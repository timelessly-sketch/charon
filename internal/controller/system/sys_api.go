package system

import (
	"charon/api/system"
	"charon/internal/service"
	"context"
)

var (
	Api = cApi{}
)

type cApi struct{}

func (c *cApi) List(ctx context.Context, _ *system.ApiListReq) (res *system.ApiListRes, err error) {
	records, err := service.SysApi().List(ctx)
	if err != nil {
		return
	}
	res = &system.ApiListRes{
		Records: records,
	}
	return
}

func (c *cApi) Edit(ctx context.Context, req *system.ApiEditReq) (_ *system.ApiListRes, err error) {
	if err = service.SysApi().Edit(ctx, req.Api); err != nil {
		return
	}
	return
}
