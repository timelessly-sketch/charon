package system

import (
	"charon/api/system"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) List(ctx context.Context, _ *system.RoleListReq) (res *system.RoleListRes, err error) {
	records, err := service.System().RoleList(ctx)
	if err != nil {
		g.Log().Warning(ctx, gerror.NewCode(gcode.CodeInternalError))
		return
	}
	res = &system.RoleListRes{Records: records}
	return
}
