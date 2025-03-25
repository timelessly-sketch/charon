package system

import (
	"charon/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"charon/api/system/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	records, total, err := service.User().List(ctx, req.UserName, req.Name, req.Page, req.Size)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError)
	}
	res = &v1.UserListRes{
		Records: records,
		Total:   total,
	}
	return
}
