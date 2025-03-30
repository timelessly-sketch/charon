package cmd

import (
	"charon/internal/router"
	"charon/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			if err := service.Config().InitConfig(ctx); err != nil {
				panic(err)
			}

			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				service.Middleware().CORS,
				service.Middleware().AuthMiddleware,
			}...)

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				router.Auth(ctx, group)
				router.System(ctx, group)
			})
			s.Run()
			return nil
		},
	}
)
