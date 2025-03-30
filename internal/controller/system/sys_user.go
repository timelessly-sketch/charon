package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) List(ctx context.Context, req *system.UserListReq) (res *system.UserListRes, err error) {
	records, total, err := service.User().List(ctx, req.UserName, req.Name, req.Page, req.Size)
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
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.User().Edit(ctx, req.User); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeDbOperationError)
	}
	return
}

func (c *cUser) Add(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error) {
	if _, err := service.User().Select(ctx, entity.User{UserName: req.UserName}); !gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeUserExists)
	}
	req.UpdatedBy = service.Middleware().GetCtxUser(ctx).UserName
	if err := service.User().Create(ctx, req.User); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeDbOperationError)
	}
	return
}
