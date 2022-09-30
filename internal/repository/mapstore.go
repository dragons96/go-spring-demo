package repository

import (
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"go-spring-demo/internal/model"
)

type MapStore struct {
	BreadStore map[uint64]*model.Bread
}

func NewMapStore() *MapStore {
	return &MapStore{
		BreadStore: make(map[uint64]*model.Bread),
	}
}

func init() {
	gs.Provide(NewMapStore).Init(func(s *MapStore) {
		log.GetLogger("MapStore").Infof("MapStore initialized successfully")
	})
}
