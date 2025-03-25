package user

import (
	"charon/internal/consts"
	"charon/internal/library/jwt"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"charon/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := service.User().UserByUserName(ctx, req.Username)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeUserNotFound)
	}
	generateJWT, err := jwt.GenerateJWT(user.Id, user.UserName)
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, gerror.NewCode(consts.CodeGenerateToken)
	}
	res = &v1.LoginRes{
		Id:       user.Id,
		Username: user.UserName,
		Avatar:   user.AvatarUrl,
		Token:    generateJWT,
		Role:     user.RoleName,
	}
	return
}
