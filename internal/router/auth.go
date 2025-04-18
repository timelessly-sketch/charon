package router

import (
	"charon/internal/controller/system"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Auth(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/auth", func(group *ghttp.RouterGroup) {
		group.Bind(system.Login)
	})
}
