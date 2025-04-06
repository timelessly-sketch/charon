package system

import (
	"charon/internal/dao"
	"charon/internal/library/cache"
	"charon/internal/library/chorm"
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	if err := sys.VerifyUnique(ctx, &public.VerifyUnique{
		Id: menu.Id,
		Where: g.Map{
			dao.Menu.Columns().Name:          menu.Name,
			dao.Menu.Columns().Path:          menu.Path,
			dao.Menu.Columns().ComponentPath: menu.ComponentPath,
		},
	}); err != nil {
		return err
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if menu.Id > 0 {
			_, err = dao.Menu.Ctx(ctx).WherePri(menu.Id).Data(&menu).Update()
			return
		}
		_, err = dao.Menu.Ctx(ctx).Insert(&menu)
		return
	}); err != nil {
		return err
	}
	_ = service.Config().LoadAuthMenu(ctx)
	return
}

// VerifyUnique 验证菜单唯一属性
func (sys *sSysMenu) VerifyUnique(ctx context.Context, in *public.VerifyUnique) (err error) {
	if in.Where == nil {
		return
	}
	cols := dao.Menu.Columns()
	msgMap := g.MapStrStr{
		cols.Path:          "菜单路径已存在，请更换",
		cols.ComponentPath: "菜单组件路径已存在，请更换",
		cols.Name:          "菜单名已经存在，请更换",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = chorm.IsUnique(ctx, &dao.Menu, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}
