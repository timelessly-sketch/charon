package system

import (
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"charon/api/system/v1"
)

func (c *ControllerV1) UserAdd(ctx context.Context, req *v1.UserAddReq) (_ *v1.UserAddRes, err error) {
	if _, err := service.User().Select(ctx, entity.User{UserName: req.UserName}); !gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeDbOperationError)
	}
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	return nil, service.User().Create(ctx, req.User)
}
