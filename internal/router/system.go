package router

import (
	"charon/internal/controller/system"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func System(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			system.Role,
			system.User,
			system.Menu,
			system.Api,
		)
	})
}
