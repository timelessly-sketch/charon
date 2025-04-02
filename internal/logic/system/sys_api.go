package system

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
)

type sSysApi struct{}

func NewSysApi() *sSysApi {
	return &sSysApi{}
}

func init() {
	service.RegisterSysApi(NewSysApi())
}

func (sys *sSysApi) List(ctx context.Context) (records []entity.Api, err error) {
	err = dao.Api.Ctx(ctx).Scan(&records)
	return
}

func (sys *sSysApi) Edit(ctx context.Context, api entity.Api) (err error) {
	if api.Id == 0 {
		if _, err = dao.Api.Ctx(ctx).Insert(&api); err != nil {
			return err
		}
	} else {
		if _, err = dao.Api.Ctx(ctx).WherePri(api.Id).Data(&api).Update(); err != nil {
			return err
		}
	}
	_ = service.Config().LoadAuthApiPath(ctx)
	return
}
