package service

import (
	"context"
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"go-spring-demo/internal/consts"
	"go-spring-demo/internal/dao"
	"go-spring-demo/internal/model"
)

func init() {
	gs.Provide(NewBreadServiceImpl).Init(func(s BreadService) {
		log.GetLogger("BreadServiceImpl").Infof("BreadService initialized successfully")
	})
}

type BreadService interface {
	Save(ctx context.Context, bread *model.Bread) (*model.Bread, error)
	QueryAll(ctx context.Context) ([]*model.Bread, error)
	QueryByName(ctx context.Context, name string) (*model.Bread, error)
	QueryById(ctx context.Context, id uint64) (*model.Bread, error)
}

type BreadServiceImpl struct {
	dao *dao.BreadDao
}

func NewBreadServiceImpl(dao *dao.BreadDao) BreadService {
	return &BreadServiceImpl{dao: dao}
}

func (s *BreadServiceImpl) Save(ctx context.Context, bread *model.Bread) (*model.Bread, error) {
	if bread.Id == 0 {
		return s.dao.Insert(ctx, bread)
	}

	_, err := s.dao.FindById(ctx, bread.Id)
	if err != nil {
		if _, ok := err.(*consts.CannotFoundError); ok {
			return s.dao.Insert(ctx, bread)
		}
		return nil, err
	}

	return s.dao.Update(ctx, bread)
}

func (s *BreadServiceImpl) QueryAll(ctx context.Context) ([]*model.Bread, error) {
	return s.dao.FindAll(ctx)
}

func (s *BreadServiceImpl) QueryByName(ctx context.Context, name string) (*model.Bread, error) {
	return s.dao.FindByName(ctx, name)
}

func (s *BreadServiceImpl) QueryById(ctx context.Context, id uint64) (*model.Bread, error) {
	return s.dao.FindById(ctx, id)
}
