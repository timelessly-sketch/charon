package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Login = sLogin{}
)

type sLogin struct{}

func (s *sLogin) Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error) {
	info, err := service.SysUser().Login(ctx, req.LoginInput)
	res = &system.LoginRes{
		LoginInfo: info,
	}
	return
}

func (s *sLogin) Logout(ctx context.Context, _ *system.LogoutReq) (_ *system.LogoutRes, err error) {
	userInfo := service.Middleware().GetCtxUser(ctx)

	key := cache.BuildUserToken(userInfo.UserName)
	if _, err := cache.Instance().Remove(ctx, key); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeCleanupTokenError)
	}
	return
}

func (s *sLogin) ResetPassword(ctx context.Context, req *system.RestPasswordReq) (_ *system.RestPasswordRes, err error) {
	err = service.SysUser().ResetPassword(ctx, req.Id)
	return
}
