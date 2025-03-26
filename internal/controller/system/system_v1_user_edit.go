package system

import (
	"charon/api/system/v1"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) UserEdit(ctx context.Context, req *v1.UserEditReq) (_ *v1.UserEditRes, err error) {
	if _, err := service.User().Select(ctx, entity.User{Id: req.Id}); gerror.Is(err, sql.ErrNoRows) {
		return nil, gerror.NewCode(gcode.CodeMissingParameter)
	}
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	return nil, service.User().Edit(ctx, req.User)
}
