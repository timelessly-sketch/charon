package system

import (
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"charon/api/system/v1"
)

func (c *ControllerV1) MenuList(ctx context.Context, _ *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	records, err := service.System().MenuList(ctx)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeDbOperationError)
	}
	res = &v1.MenuListRes{
		Records: records,
	}
	return
}
