# Go-Spring Demo 项目

#### 手把手带你从0构建一个完整的Go-Spring项目

### Go-Spring 官方地址: [https://github.com/go-spring/go-spring](https://github.com/go-spring/go-spring) 

## 起步
### 快速开始(最小化应用)
1. 创建一个项目目录, 这里以go-spring-demo-simple目录为例, 并在目录下创建一个main.go文件, 目录结构如下：
```text
go-spring-demo-simple
└- main.go
```
2. 安装go-spring依赖：
```shell
go get -u github.com/go-spring/spring-core@v1.1.1
# gin web服务器
go get -u github.com/go-spring/starter-gin@v1.1.1
# echo web服务器, 与gin web服务器二选一即可
go get -u github.com/go-spring/starter-echo@v1.1.1
```
3. 编辑main.go文件:

```go
package main

import (
	"context"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	// 添加gin web服务starter
	_ "github.com/go-spring/starter-gin"
	// 添加echo web服务starter, 与gin web服务二选一即可
	//_ "github.com/go-spring/starter-echo"
	"log"
)

func init() {
	// 创建一个路由为/get-mapping的get请求资源, gin风格
	gs.GetMapping("/get-mapping", func(ctx web.Context) {
		ctx.JSON(web.SUCCESS.Data("get-mapping"))
	})
	// 创建一个路由为/get-binding的get请求资源, 对象风格
	gs.GetBinding("/get-binding", func(ctx context.Context, req *GetBindingReq) *GetBindingResp {
		return &GetBindingResp{Say: "hello " + req.Name}
	})
}

type GetBindingReq struct {
	Name string `form:"name" json:"name"`
}

type GetBindingResp struct {
	Say string `json:"say"`
}

func main() {
	log.Fatalln(gs.Run())
}
```
4. 到这里就完成了一个简单的web服务, 运行`main.go`, 访问 `http://127.0.0.1:8080/get-mapping` 和 `http://127.0.0.1:8080/get-binding?name=dragons` 即可观察输出结果


### 更多内容见: [完整项目搭建示例](docs/zh/full_project/outline.md)