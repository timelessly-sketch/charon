package user

import (
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"charon/api/user/v1"
)

func (c *ControllerV1) Info(ctx context.Context, _ *v1.InfoReq) (res *v1.InfoRes, err error) {
	name := g.RequestFromCtx(ctx).GetCtxVar("name").String()
	user, _ := service.User().UserByUserName(ctx, name)
	res = &v1.InfoRes{
		Id:       user.Id,
		NickName: user.NickName,
		Roles:    gconv.Strings(user.RoleName),
	}
	return
}
