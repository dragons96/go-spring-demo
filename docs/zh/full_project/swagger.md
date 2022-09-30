## Swagger工具集成
### 1. 安装依赖:
```shell
# swag cli工具
go install github.com/swaggo/swag/cmd/swag@latest

# swagger依赖
go get -u github.com/swaggo/http-swagger@v1.3.3
go get -u github.com/swaggo/swag@v1.8.6
```
### 2. 修改`internal/controller/bread.go`文件, 增加控制器方法[swagger注释](https://github.com/swaggo/swag), 修改后文件如下:
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

```

### 3. 修改项目根目录`main.go`文件, 增加基本信息[swagger注释](https://github.com/swaggo/swag), 修改后文件内容如下:
```go
package main

import (
	// app
	"go-spring-demo/internal/cmd"
)

// @title Bread Server
// @version 0.0.1
// @description This is a go-spring demo.
// @termsOfService https://github.com/go-spring/go-spring
// @contact.name github.com/go-spring/go-spring
// @contact.url https://github.com/go-spring/go-spring
// @contact.email 521274311@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @basePath /api/v1/bread
func main() {
	cmd.Execute()
}

```

### 4. 在`gen/v1`目录下新建swagger.bat(windows系统)或swagger.sh(linux系统)文件, 编写swagger生成脚本如下:
`swagger.bat`
```shell
set pathprefix=这里填项目绝对路径

swag init -g %pathprefix%/main.go -o %pathprefix%/api/v1 --pd
```
`swagger.sh`
```shell
pathprefix=这里填项目绝对路径

swag init -g ${pathprefix}/main.go -o ${pathprefix}/api/v1 --pd
```
### 5. 运行`gen/v1/swagger.bat`或`gen/v1/swagger.sh`脚本, 若在`api/v1`目录下生成`docs.go`, `swagger.json`, `swagger.yaml`文件则说明脚本运行成功
### 6. 在`api/v1`目录下新建`swagger.go`文件, 内如如下:
```go
package v1

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Init() {
	gs.GetMapping("/swagger/*", func(ctx web.Context) {
		hs := httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json"))
		hs.ServeHTTP(ctx.Response(), ctx.Request())
	})
}
```
### 这里简单说明一下什么情况下使用`init`方法, 什么情况下使用`Init`方法. 一般情况下如果希望对象自动注入spring容器, 则使用`init`方法, 若希望自己通过条件或其他方式控制注入容器, 则使用`Init`方法, 这里我们会在cli工具中添加swagger启动参数, 所以选择`Init`方法

### 7. 修改`internal/cmd/bread.go`文件, 修改后内容如下:
```go
package cmd

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/spf13/cobra"
	v1 "go-spring-demo/api/v1"
	"log"
	"os"
)

var (
	// 定义启动变量
	swagger    bool
)

func init() {
	// 添加参数绑定
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
		// 条件注入
		if swagger {
			v1.Init()
		}
		log.Fatalln(gs.Run())
	},
}
```
### 8. 到这里swagger就已经配置完了, 对比以下两个命令, 有`-s`参数的命令将会开启swagger文档, swagger地址：http://127.0.0.1:8080/swagger
```shell
go run main.go

go run main.go -s
```

### 接下来我们继续改造添加grpc服务

#### [上一页：cli工具](cli.md)

#### [下一页：grpc集成](grpc.md)