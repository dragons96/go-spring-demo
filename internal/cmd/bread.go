package cmd

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/starter-grpc/server/factory"
	"github.com/spf13/cobra"
	v1 "go-spring-demo/api/v1"
	"go-spring-demo/internal/controller/adapter"
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
			adapter.Init()
			// grpc server 启动
			gs.Provide(factory.NewStarter, "${grpc.server}").Export((*gs.AppEvent)(nil))
		}
		log.Fatalln(gs.Run())
	},
}
