package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

func (c *cMenu) List(ctx context.Context, req *system.MenuListReq) (res *system.MenuListRes, err error) {
	getRecords := func() ([]entity.Menu, error) {
		if req.Id == 0 {
			return service.System().MenuList(ctx)
		}
		return service.System().MenuDynamic(ctx, req.Id)
	}

	records, err := getRecords()
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	res = &system.MenuListRes{
		Records: records,
	}
	return
}

func (c *cMenu) Edit(ctx context.Context, req *system.MenuEditReq) (_ *system.MenuEditRes, err error) {
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().MenuEdit(ctx, req.Menu); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	return
}

func (c *cMenu) Add(ctx context.Context, req *system.MenuAddReq) (res *system.MenuAddRes, err error) {
	if _, err := service.System().MenuByName(ctx, req.Name); !gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeMenuExists)
	}

	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().MenuAdd(ctx, req.Menu); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	return
}
