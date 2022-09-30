package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"go-spring-demo/internal/consts"
	"go-spring-demo/internal/model"
	"go-spring-demo/internal/repository"
	"sync"
	"sync/atomic"
)

func init() {
	gs.Provide(NewBreadDao).Init(func(dao *BreadDao) {
		log.GetLogger("BreadDao").Infof("BreadDao initialized successfully")
	})
}

type BreadDao struct {
	store *repository.MapStore
	mu    sync.RWMutex
	pri   uint64
}

func NewBreadDao(store *repository.MapStore) *BreadDao {
	return &BreadDao{store: store}
}

func (d *BreadDao) Insert(ctx context.Context, bread *model.Bread) (*model.Bread, error) {
	_, err := d.FindByName(ctx, bread.Name)
	if err == nil {
		return nil, errors.New(fmt.Sprintf("Cannot create Bread because it already exists with that name %q", bread.Name))
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	bread.Id = atomic.AddUint64(&d.pri, 1)
	d.store.BreadStore[bread.Id] = bread
	return bread, nil
}

func (d *BreadDao) Update(ctx context.Context, bread *model.Bread) (*model.Bread, error) {
	d.mu.RLock()
	oh, ok := d.store.BreadStore[bread.Id]
	d.mu.RUnlock()
	if !ok {
		return nil, consts.CannotFoundIdError.Value(bread.Id)
	}

	d.mu.Lock()
	d.store.BreadStore[bread.Id] = bread
	d.mu.Unlock()

	return oh, nil
}

func (d *BreadDao) DeleteById(ctx context.Context, id uint64) (*model.Bread, error) {
	d.mu.RLock()
	oh, ok := d.store.BreadStore[id]
	d.mu.RUnlock()
	if !ok {
		return nil, consts.CannotFoundIdError.Value(id)
	}

	d.mu.Lock()
	delete(d.store.BreadStore, id)
	d.mu.Unlock()

	return oh, nil
}

func (d *BreadDao) FindAll(ctx context.Context) ([]*model.Bread, error) {
	ans := make([]*model.Bread, 0, len(d.store.BreadStore))
	for _, v := range d.store.BreadStore {
		ans = append(ans, v)
	}
	return ans, nil
}

func (d *BreadDao) FindById(ctx context.Context, id uint64) (*model.Bread, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	h, ok := d.store.BreadStore[id]
	if !ok {
		return nil, consts.CannotFoundIdError.Value(id)
	}

	return h, nil
}

func (d *BreadDao) FindByName(ctx context.Context, name string) (*model.Bread, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	breads, err := d.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, bread := range breads {
		if bread.Name == name {
			return bread, nil
		}
	}

	return nil, consts.CannotFoundNameError.Value(name)
}
