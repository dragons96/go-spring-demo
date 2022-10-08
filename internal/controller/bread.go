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

func InitBread() {
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
// @summary saves bread
// @title Bread.Save
// @tags Bread.Save
// @router / [post]
// @param req body SaveReq true "Save Params"
// @success 200 {object} web.RpcResult
// @failure 404 {object} string
// @failure 200 {object} web.RpcResult
func (c *BreadController) Save(ctx context.Context, req *SaveReq) *web.RpcResult {
	bread := model.NewBread(req.Id, req.Name)
	return c.wrapperDataError(c.service.Save(ctx, bread))
}

// QueryAll is query all bread method
// @summary queries all bread
// @title Bread.QueryAll
// @tags Bread.QueryAll
// @router / [get]
// @success 200 {object} web.RpcResult
// @failure 404 {object} string
// @failure 200 {object} web.RpcResult
func (c *BreadController) QueryAll(ctx context.Context, req *QueryAllReq) *web.RpcResult {
	return c.wrapperDataError(c.service.QueryAll(ctx))
}

// Query is query some bread method
// @summary queries some bread
// @title Bread.Query
// @tags Bread.Query
// @router /id/{id} [get]
// @router /name/{name} [get]
// @param req body QueryReq true "Query Params"
// @success 200 {object} web.RpcResult
// @failure 404 {object} string
// @failure 200 {object} web.RpcResult
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
