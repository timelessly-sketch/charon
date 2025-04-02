package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) List(ctx context.Context, _ *system.RoleListReq) (res *system.RoleListRes, err error) {
	records, err := service.SysMenu().RoleList(ctx)
	if err != nil {
		g.Log().Warning(ctx, gerror.NewCode(consts.CodeDbOperationError))
		return
	}
	res = &system.RoleListRes{Records: records}
	return
}
