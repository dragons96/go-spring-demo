package main

import (
	"context"
	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-spring/starter-grpc/client"
	__ "go-spring-demo/internal/proto/pb"
	"log"
)

func main() {
	// gs.GrpcClient(..., endpoint) 中的 endpoint 需要与配置文件中的名称一一对应
	gs.GrpcClient(__.NewBreadClient, "bread").Init(func(client __.BreadClient) {
		data, err := client.Save(context.TODO(), &__.SaveBreadReq{
			Name: "红豆面包",
		})

		if err != nil {
			log.Printf("错误: %v\n", err)
		}

		data, err = client.Query(context.TODO(), &__.QueryBreadReq{Query: &__.QueryBreadReq_Id{Id: 1}})
		if err != nil {
			log.Printf("错误: %v\n", err)
		}

		log.Printf("查询ID为1的面包: %v\n", data)
	})
	log.Fatalln(gs.Web(false).Run())
}
