package system

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
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
	err = dao.Menu.Ctx(ctx).WhereNot(dao.Menu.Columns().Id, 0).Scan(&records)
	return
}

func (sys *sSystem) MenuByName(ctx context.Context, name string) (record entity.Menu, err error) {
	err = dao.Menu.Ctx(ctx).Where("name = ?", name).Scan(&record)
	return
}

func (sys *sSystem) MenuCreate(ctx context.Context, menu entity.Menu) (err error) {
	_, err = dao.Menu.Ctx(ctx).OmitEmpty().Insert(&menu)
	return
}

func (sys *sSystem) MenuEdit(ctx context.Context, menu entity.Menu) (err error) {
	_, err = dao.Menu.Ctx(ctx).Where(dao.User.Columns().Id, menu.Id).Data(&menu).Update()
	return
}
