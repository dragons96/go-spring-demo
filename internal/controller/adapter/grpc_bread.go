package adapter

import (
	"context"
	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/gs"
	"go-spring-demo/internal/model"
	__ "go-spring-demo/internal/proto/pb"
	"go-spring-demo/internal/service"
	"google.golang.org/protobuf/types/known/anypb"
)

func Init() {
	gs.Provide(NewGrpcBreadServer).Init(func(s __.BreadServer) {
		gs.GrpcServer("bread.Bread", &grpc.Server{
			Register: __.RegisterBreadServer,
			Service:  s,
		})
	})
}

type breadServerAdapterController struct {
	*__.UnimplementedBreadServer
	service service.BreadService
}

func NewGrpcBreadServer(s service.BreadService) __.BreadServer {
	return &breadServerAdapterController{service: s}
}

func (b *breadServerAdapterController) Save(ctx context.Context, req *__.SaveBreadReq) (*__.CommonOneResp, error) {
	bread, err := b.service.Save(ctx, model.NewBread(req.Id, req.Name))
	if err != nil {
		return nil, err
	}

	data, err := anypb.New(&__.BreadStore{
		Id:   bread.Id,
		Name: bread.Name,
	})
	if err != nil {
		return nil, err
	}

	return &__.CommonOneResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}, nil
}

func (b *breadServerAdapterController) QueryAll(ctx context.Context, req *__.QueryAllBreadReq) (*__.CommonManyResp, error) {
	breads, err := b.service.QueryAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := make([]*anypb.Any, 0, len(breads))
	for _, bread := range breads {
		one, err := anypb.New(&__.BreadStore{
			Id:   bread.Id,
			Name: bread.Name,
		})
		if err != nil {
			return nil, err
		}
		tmp = append(tmp, one)
	}
	return &__.CommonManyResp{
		Code: 0,
		Msg:  "success",
		Data: tmp,
	}, nil
}

func (b *breadServerAdapterController) Query(ctx context.Context, req *__.QueryBreadReq) (*__.CommonOneResp, error) {
	var (
		bread *model.Bread
		err   error
	)

	switch req.Query.(type) {
	case *__.QueryBreadReq_Id:
		bread, err = b.service.QueryById(ctx, req.Query.(*__.QueryBreadReq_Id).Id)
	case *__.QueryBreadReq_Name:
		bread, err = b.service.QueryByName(ctx, req.Query.(*__.QueryBreadReq_Name).Name)
	}
	if err != nil {
		return nil, err
	}

	data, err := anypb.New(&__.BreadStore{
		Id:   bread.Id,
		Name: bread.Name,
	})
	if err != nil {
		return nil, err
	}

	return &__.CommonOneResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}, nil
}
