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
	User = cUser{}
)

type cUser struct{}

func (c *cUser) List(ctx context.Context, req *system.UserListReq) (res *system.UserListRes, err error) {
	records, total, err := service.SysUser().List(ctx, req.UserName, req.Name, req.Page, req.Size)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeDbOperationError)
	}

	res = &system.UserListRes{
		Records: records,
		Total:   total,
	}
	return
}

func (c *cUser) Edit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	if err := service.SysUser().Edit(ctx, &req.User); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeDbOperationError)
	}
	return
}
