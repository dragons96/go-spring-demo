## 这里以一个面包管理服务为例。（为了方便起见, 仅保存ID以及名称字段, 且使用内存map结构模拟数据库）  
### 1. 下载对应系统的 [protoc工具](https://github.com/protocolbuffers/protobuf/releases) 并添加到系统环境变量, 打开命令行执行命令 `protoc --help`, 若正常显示帮助详情则说明安装成功    
### 2. 安装依赖安装依赖, 命令如下:
```shell
# go spring 基础依赖
go get -u github.com/go-spring/spring-base@v1.1.1
go get -u github.com/go-spring/spring-core@v1.1.1
go get -u github.com/go-spring/starter-gin@v1.1.1 或者 go get -u github.com/go-spring/starter-echo@v1.1.1
```
### 3. 在`internal/model`中新建`bread.go`文件, 内容如下:
```go
package model

type Bread struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewBread(id uint64, name string) *Bread {
	return &Bread{Id: id, Name: name}
}

func (h *Bread) String() string {
	return "Bread: " + h.Name
}
```
### 4. 在`internal/repository`中新建`mapstore.go`文件, 内容如下:
```go
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
	// gs.Provide 为函数式注入/构造方法注入, 函数的参数会从spring容器中获取,函数的返回值会注入到spring容器中
	// gs.Object 为对象实例注入, 该方法可改写为 gs.Object(NewMapStore()).Init(...)
	gs.Provide(NewMapStore).Init(func(s *MapStore) {
		log.GetLogger("MapStore").Infof("MapStore initialized successfully")
	})
}
```
### 5. 在`internal/dao`中新建bread.go文件, 内容如下:
```go
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
```
### 6. 在`internal/service`中新建bread.go文件, 内容如下:
```go
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
```
### 7. 在`internal/controller`中新建bread.go文件, 内容如下:
```go
package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	"go-spring-demo/internal/model"
	"go-spring-demo/internal/service"
)

func init() {
	gs.Provide(NewBreadController).Init(func(c *BreadController) {
		log.GetLogger("BreadController").Infof("BreadController initialized successfully")
		gs.PostBinding("/api/v1/bread", c.Save)
		gs.GetBinding("/api/v1/bread", c.QueryAll)
		gs.GetBinding("/api/v1/bread/id/:id", c.Query)
		gs.GetBinding("/api/v1/bread/name/:name", c.Query)
	})
}

type BreadController struct {
	service service.BreadService
}

func NewBreadController(service service.BreadService) *BreadController {
	return &BreadController{service: service}
}

// Save is create or update a bread method
func (c *BreadController) Save(ctx context.Context, req *SaveReq) *web.RpcResult {
	bread := model.NewBread(req.Id, req.Name)
	return c.wrapperDataError(c.service.Save(ctx, bread))
}

// QueryAll is query all bread method
func (c *BreadController) QueryAll(ctx context.Context, req *QueryAllReq) *web.RpcResult {
	return c.wrapperDataError(c.service.QueryAll(ctx))
}

// Query is query some bread method
func (c *BreadController) Query(ctx context.Context, req *QueryReq) *web.RpcResult {
	if req.Id > 0 {
		return c.wrapperDataError(c.service.QueryById(ctx, req.Id))
	}

	if req.Name != "" {
		return c.wrapperDataError(c.service.QueryByName(ctx, req.Name))
	}
	
	return web.ERROR.Error(errors.New(fmt.Sprintf("Cannot find query parameter \"id\" and \"name\"")))
}

func (c *BreadController) wrapperDataError(data interface{}, err error) *web.RpcResult {
	if err != nil {
		return web.ERROR.Error(err)
	}
	return web.SUCCESS.Data(data)
}

type QueryAllReq struct{}

type QueryReq struct {
	Id   uint64 `uri:"id"`
	Name string `uri:"name"`
}

type SaveReq struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
```
### 8. 在项目根目录创建`main.go`文件, 内容如下:
```go
package main

import (
	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-spring/starter-gin"
	_ "go-spring-demo/internal/controller"
	"log"
)

func main() {
	log.Fatalln(gs.Run())
}
```
### 9. 运行`main.go`, 启动程序:
```shell
go run main.go
```
### 10. 到这里一个完整的web项目已经完成了, 但是功能还很简陋, 接下来我们来改造cli工具


#### [上一页：项目结构](structure.md) 

#### [下一页：cli工具](cli.md)