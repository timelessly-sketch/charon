package router

import (
	"charon/internal/controller/cluster"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Service(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/service", func(group *ghttp.RouterGroup) {
		group.Bind(
			cluster.Cluster,
		)
	})
}
