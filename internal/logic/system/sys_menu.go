package system

import (
	"charon/internal/consts"
	"charon/internal/dao"
	"charon/internal/library/cache"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysMenu struct{}

func NewSysMenu() *sSysMenu {
	return &sSysMenu{}
}

func init() {
	service.RegisterSysMenu(NewSysMenu())
}

func (sys *sSysMenu) RoleList(ctx context.Context) (records []entity.Role, err error) {
	if err = dao.Role.Ctx(ctx).Scan(&records); err != nil {
		return nil, err
	}
	return
}

func (sys *sSysMenu) MenuList(ctx context.Context, id int) (records []entity.Menu, err error) {
	if id == 0 {
		if err = dao.Menu.Ctx(ctx).Scan(&records); err != nil {
			return nil, err
		}
		return
	}
	role := service.Middleware().GetCtxUser(ctx).RoleName
	value, err := cache.Instance().Get(ctx, cache.BuildMenu(role))
	if err = gconv.Structs(value, &records); err != nil {
		return nil, err
	}
	return
}

func (sys *sSysMenu) MenuEdit(ctx context.Context, menu *entity.Menu) (err error) {
	if count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Path, menu.Path).Count(); count > 0 || err != nil {
		return gerror.NewCode(consts.CodeMenuExists)
	}
	if menu.Id > 0 {
		if _, err = dao.Menu.Ctx(ctx).WherePri(menu.Id).Data(&menu).Update(); err != nil {
			return err
		}
		return
	}
	_, err = dao.Menu.Ctx(ctx).Insert(&menu)
	return
}
