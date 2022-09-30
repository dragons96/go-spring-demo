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
