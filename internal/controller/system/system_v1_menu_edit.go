package system

import (
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"charon/api/system/v1"
)

func (c *ControllerV1) MenuEdit(ctx context.Context, req *v1.MenuEditReq) (_ *v1.MenuEditRes, err error) {
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().MenuEdit(ctx, req.Menu); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeDbOperationError)
	}
	return
}
