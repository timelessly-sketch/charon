package system

import (
	v1 "charon/api/system/v1"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	records, err := service.System().RoleList(ctx)
	if err != nil {
		g.Log().Warning(ctx, gerror.NewCode(gcode.CodeDbOperationError))
		return
	}
	res = &v1.RoleListRes{Records: records}
	return
}
