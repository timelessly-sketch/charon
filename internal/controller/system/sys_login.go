package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/library/token"
	"charon/internal/model/entity"
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
	user, err := service.User().Select(ctx, entity.User{UserName: req.Username})
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeUserNotFound)
	}
	generateJWT, err := token.GenerateJWT(ctx, user)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeGenerateTokenError)
	}

	res = &system.LoginRes{
		Id:       user.Id,
		Username: user.UserName,
		Avatar:   user.AvatarUrl,
		Token:    generateJWT,
		Role:     user.RoleName,
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
