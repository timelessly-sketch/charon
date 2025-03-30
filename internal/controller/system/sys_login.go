package system

import (
	"charon/api/system"
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/library/token"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Login = sLogin{}
)

type sLogin struct{}

func (s *sLogin) Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error) {
	user, err := service.System().UserSelect(ctx, entity.User{UserName: req.Username})
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeUserNotFound)
	}
	if s.generatePassword(req.Password) != user.Password {
		return nil, gerror.NewCode(consts.CodePasswordInvalid)
	}

	generateJWT, err := token.GenerateJWT(ctx, user)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, err
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

func (s *sLogin) ResetPassword(ctx context.Context, req *system.RestPasswordReq) (_ *system.RestPasswordRes, err error) {
	user, err := service.System().UserSelect(ctx, entity.User{Id: req.Id})
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeUserNotFound)
	}

	user.Password = s.generatePassword(user.UserName)
	if err := service.System().UserUpdate(ctx, user); err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeResetPassError)
	}
	return
}

func (s *sLogin) generatePassword(password string) string {
	hash := sha256.New()
	hash.Write(gconv.Bytes(password))
	return hex.EncodeToString(hash.Sum(nil))
}
