package cluster

import (
	"charon/api/cluster"
	"charon/internal/consts"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Cluster = cCluster{}
)

type cCluster struct{}

func (c *cCluster) List(ctx context.Context, req *cluster.ListReq) (res *cluster.ListRes, err error) {
	records, total, err := service.Cluster().List(ctx, req.Page, req.Size)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeDbOperationError)
	}
	res = &cluster.ListRes{
		Records: records,
		Total:   total,
	}
	return
}
