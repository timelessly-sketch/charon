package system

import (
	"charon/api/system"
	"charon/internal/service"
	"context"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

func (c *cMenu) List(ctx context.Context, req *system.MenuListReq) (res *system.MenuListRes, err error) {
	records, err := service.SysMenu().MenuList(ctx, req.Id)
	res = &system.MenuListRes{
		Records: records,
	}
	return
}

func (c *cMenu) Edit(ctx context.Context, req *system.MenuEditReq) (_ *system.MenuEditRes, err error) {
	if err = service.SysMenu().MenuEdit(ctx, &req.Menu); err != nil {
		return
	}
	_ = service.Config().LoadAuthMenu(ctx)
	return
}
