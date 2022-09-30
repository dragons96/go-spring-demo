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
