## 在上文中我们已经完成面包的增删改查功能， 接下来我们来改造cli工具
### 1. 安装依赖:
```shell
# 命令行工具依赖
go get -u github.com/spf13/cobra@v1.5.0
```
### 2. 在`internal/cmd`目录下新建`cmd.go`文件, 内容如下:
```go
package cmd

// 该文件仅配置需要自动注入的包即可, 比如 starter-gin, 和 controller
import (
	_ "github.com/go-spring/starter-gin"
	_ "go-spring-demo/internal/controller"
)
```
### 3. 在`internal/cmd`目录下新建`bread.go`文件, 内容如下:
```go
package cmd

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/spf13/cobra"
	"log"
	"os"
)

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
		log.Fatalln(gs.Run())
	},
}
```
### 4. 修改根目录`main.go`文件, 内容如下:
```go
import (
	// app
	"go-spring-demo/internal/cmd"
)

func main() {
	cmd.Execute()
}
```
### 运行`main.go`即可, 到这里cli工具已经建设好了, 目前尚未配置参数可能看不出效果, 后文中会一步步演示如何使用cli工具, 接下来我们继续改造添加swagger

#### [上一页：项目基础功能](base.md)

#### [下一页：swagger集成](swagger.md)