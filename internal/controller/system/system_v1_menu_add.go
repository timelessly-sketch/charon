package system

import (
	"charon/internal/consts"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"charon/api/system/v1"
)

func (c *ControllerV1) MenuAdd(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error) {
	if _, err := service.System().MenuByName(ctx, req.Name); !gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeMenuExists)
	}
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().MenuCreate(ctx, req.Menu); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	return
}
