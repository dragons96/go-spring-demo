## Grpc 集成
### 1. 安装依赖
```shell
# grpc cli工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# grpc 依赖
go get -u github.com/go-spring/starter-grpc@v1.1.1
go get -u google.golang.org/grpc@v1.49.0
go get -u google.golang.org/protobuf@v1.28.1
```
### 2. 在`internal/proto`目录下新建`common.proto`和`bread.proto`文件, 用来编写grpc传输对象, 文件内容如下:
`common.proto`
```protobuf
syntax = "proto3";

option go_package = ".";

package common;

import "google/protobuf/any.proto";

/**
{
  "code": 0,
  "msg": "success",
  "data": {}
}
 */
message CommonOneResp {
  uint32 code = 1;
  string msg = 2;
  google.protobuf.Any data = 3;
}

/**
{
  "code": 0,
  "msg": "success",
  "data": []
}
 */
message CommonManyResp {
  uint32 code = 1;
  string msg = 2;
  repeated google.protobuf.Any data = 3;
}
```
`bread.proto`
```protobuf
syntax = "proto3";

option go_package = ".";

package bread;

import "internal/proto/common.proto";

service Bread {
  rpc Save(SaveBreadReq) returns(common.CommonOneResp);
  rpc QueryAll(QueryAllBreadReq) returns(common.CommonManyResp);
  rpc Query(QueryBreadReq) returns(common.CommonOneResp);
}

message SaveBreadReq {
  uint64 id = 1;
  string name = 2;
}

message QueryBreadReq {
  oneof query {
    uint64 id = 1;
    string name = 2;
  }
}

message QueryAllBreadReq {}

message BreadStore {
  uint64 id = 1;
  string name = 2;
}
```
### 3. 在`gen`目录下创建`protobuf`目录, 并在`gen/protobuf`目录下新建`generator.bat`或`generator.sh`文件, 编写生成pb.go文件脚本内容如下:
`generator.bat`
```shell
set pathprefix=这里填项目绝对路径

protoc -I %pathprefix% --go_out=%pathprefix%/internal/proto/pb --go-grpc_out=%pathprefix%/internal/proto/pb %pathprefix%/internal/proto/*.proto
```
`generator.sh`
```shell
#!/bin/bash
pathprefix=这里填项目绝对路径

protoc -I %pathprefix% --go_out=%pathprefix%/internal/proto/pb --go-grpc_out=%pathprefix%/internal/proto/pb %pathprefix%/internal/proto/*.proto
```
运行`generator.bat`或`generator.sh`脚本, 在`internal/proto/pb`目录下将会生成`bread.pb.go`、`bread_grpc.pb.go`、`common.pb.go`文件, 则说明脚本运行正常
### 4. 在`internal/controller`目录下新建`grpc`目录, 并在`internal/controller/grpc`目录下新建`bread.go`文件，文件内容如下:
```go
package grpc

import (
	"context"
	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/gs"
	"go-spring-demo/internal/model"
	__ "go-spring-demo/internal/proto/pb"
	"go-spring-demo/internal/service"
	"google.golang.org/protobuf/types/known/anypb"
)

func InitBread() {
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
```

### 5. 修改`internal/cmd/bread.go`文件, 为grpc server添加启动命令开关, 修改后内容如下:
```go
package cmd

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/starter-grpc/server/factory"
	"github.com/spf13/cobra"
	v1 "go-spring-demo/api/v1"
	"go-spring-demo/internal/controller/grpc"
	"log"
	"os"
)

var (
	grpcServer bool
	swagger    bool
)

func init() {
	rootCmd.Flags().BoolVarP(&grpcServer, "grpc-server", "g", false, "enable grpc server")
	rootCmd.Flags().BoolVarP(&swagger, "swagger", "s", false, "enable swagger docs")
}

// Execute is a command executor
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "bread",
	Short: "Bread management server",
	Long:  "Bread management server",
	Run: func(cmd *cobra.Command, args []string) {
		if swagger {
			v1.Init()
		}
		if grpcServer {
			// grpc endpoint 注入
			grpc.InitBread()
			// grpc server 启动
			gs.Provide(factory.NewStarter, "${grpc.server}").Export((*gs.AppEvent)(nil))
		}
		log.Fatalln(gs.Run())
	},
}
```

### 6. 修改`config`目录下配置文件, 增加`grpc`配置, 内容如下:
```yaml
web:
  server:
    port: 8080

grpc:
  # grpc server 配置
  server:
    port: 8081
  # 正式环境下若无远程调用无需配endpoint, 这里为了方便配置写在同一个文件
  # grpc client 配置
  endpoint:
    # 需要访问的 grpc server 名称, 可随意定义, 代码中的endpoint需与该名称对应
    bread:
      # grpc server address
      address: localhost:8081

```
### 7. 创建`client`目录, 并创建`client/main.go`文件编写客户端测试, 文件内容如下: (非必须步骤, 这里仅为了测试使用)
```go
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
```

### 8. 启动 `go run main.go -g -s`, 也可通过命令 `go run main.go -h` 查看命令帮助

## 拓展, 为http server配置开关命令，这里以gin为例, echo同理:
### 1. 修改`internal/cmd/cmd.go`文件, 内容如下:
```go
package cmd

// 移除两个自动加载包
//import (
	//_ "github.com/go-spring/starter-gin"
	//_ "go-spring-demo/internal/controller"
//)
```
### 2. 修改`internal/controller/bread.go`文件, 将init方法名改为InitBread
### 3. 修改`internal/cmd/bread.go`文件, 内容如下:
```go
package cmd

import (
	"github.com/go-spring/spring-core/gs"
	SpringGin "github.com/go-spring/spring-gin"
	"github.com/go-spring/starter-grpc/server/factory"
	"github.com/spf13/cobra"
	v1 "go-spring-demo/api/v1"
	"go-spring-demo/internal/controller"
	"go-spring-demo/internal/controller/grpc"
	"log"
	"os"
)

var (
	webServer  bool
	grpcServer bool
	swagger    bool
)

func init() {
	rootCmd.Flags().BoolVarP(&webServer, "web-server", "w", false, "enable web server")
	rootCmd.Flags().BoolVarP(&grpcServer, "grpc-server", "g", false, "enable grpc server")
	rootCmd.Flags().BoolVarP(&swagger, "swagger", "s", false, "enable swagger docs")
}

// Execute is a command executor
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "bread",
	Short: "Bread management server",
	Long:  "Bread management server",
	Run: func(cmd *cobra.Command, args []string) {
		if swagger {
			v1.Init()
		}
		if webServer {
			// web endpoint 注入
			controller.InitBread()
			// web server 启动
			gs.Provide(SpringGin.New, "${web.server}")
		}
		if grpcServer {
			// grpc endpoint 注入
			grpc.InitBread()
			// grpc server 启动
			gs.Provide(factory.NewStarter, "${grpc.server}").Export((*gs.AppEvent)(nil))
		}
		log.Fatalln(gs.Web(webServer).Run())
	},
}
```

### 接下来我们继续改造集成gorm框架

#### [上一页：swagger集成](swagger.md)

#### [下一页：gorm集成](gorm.md)