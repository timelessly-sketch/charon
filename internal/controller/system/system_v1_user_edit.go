package system

import (
	"charon/api/system/v1"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UserEdit(ctx context.Context, req *v1.UserEditReq) (_ *v1.UserEditRes, err error) {
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.User().Edit(ctx, req.User); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeDbOperationError)
	}
	return
}
