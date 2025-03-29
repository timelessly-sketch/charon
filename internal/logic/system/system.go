package system

import (
	"charon/internal/dao"
	"charon/internal/library/cache"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystem struct{}

func NewSystem() *sSystem {
	return &sSystem{}
}

func init() {
	service.RegisterSystem(NewSystem())
}

func (sys *sSystem) RoleList(ctx context.Context) (records []entity.Role, err error) {
	err = dao.Role.Ctx(ctx).Scan(&records)
	return
}

func (sys *sSystem) MenuList(ctx context.Context) (records []entity.Menu, err error) {
	err = dao.Menu.Ctx(ctx).Scan(&records)
	return
}

func (sys *sSystem) MenuDynamic(ctx context.Context, id int) (records []entity.Menu, err error) {
	var (
		user entity.User
	)
	if err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&user); err != nil {
		return nil, err
	}
	value, err := cache.Instance().Get(ctx, user.RoleName+":Menu")
	if err = gconv.Structs(value, &records); err != nil {
		return nil, err
	}
	return
}

func (sys *sSystem) MenuByName(ctx context.Context, name string) (record entity.Menu, err error) {
	err = dao.Menu.Ctx(ctx).Where("name = ?", name).Scan(&record)
	return
}

func (sys *sSystem) MenuAdd(ctx context.Context, menu entity.Menu) (err error) {
	_, err = dao.Menu.Ctx(ctx).OmitEmpty().Insert(&menu)
	return
}

func (sys *sSystem) MenuEdit(ctx context.Context, menu entity.Menu) (err error) {
	_, err = dao.Menu.Ctx(ctx).Where(dao.User.Columns().Id, menu.Id).Data(&menu).Update()
	return
}

func (sys *sSystem) ApiList(ctx context.Context) (records []entity.Api, err error) {
	err = dao.Api.Ctx(ctx).Scan(&records)
	return
}

func (sys *sSystem) ApiEdit(ctx context.Context, api entity.Api) (err error) {
	_, err = dao.Api.Ctx(ctx).Where(dao.Api.Columns().Id, api.Id).Data(&api).Update()
	return
}

func (sys *sSystem) ApiAdd(ctx context.Context, api entity.Api) (err error) {
	_, err = dao.Api.Ctx(ctx).OmitEmpty().Insert(&api)
	return
}

func (sys *sSystem) ApiByName(ctx context.Context, name string) (api entity.Api, err error) {
	err = dao.Api.Ctx(ctx).Where("name = ?", name).Scan(&api)
	return
}
