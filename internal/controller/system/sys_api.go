package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Api = cApi{}
)

type cApi struct{}

func (c *cApi) List(ctx context.Context, _ *system.ApiListReq) (res *system.ApiListRes, err error) {
	records, err := service.System().ApiList(ctx)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	res = &system.ApiListRes{
		Records: records,
	}
	return
}

func (c *cApi) Edit(ctx context.Context, req *system.ApiEditReq) (_ *system.ApiListRes, err error) {
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().ApiEdit(ctx, req.Api); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	_ = service.Config().LoadAuthApiPath(ctx)
	return
}

func (c *cApi) Add(ctx context.Context, req *system.ApiAddReq) (_ *system.ApiListRes, err error) {
	if _, err := service.System().ApiByName(ctx, req.Name); !gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeApiExists)
	}

	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.System().ApiAdd(ctx, req.Api); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	_ = service.Config().LoadAuthApiPath(ctx)
	return
}
