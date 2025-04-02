package system

import (
	"charon/internal/model/public"
	"github.com/gogf/gf/v2/frame/g"
)

// 只提供登录接口, 用户在Ldap上自行注册或者管理员手动添加

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"用户登录" noAuth:"true"`
	public.LoginInput
}

type LoginRes struct {
	public.LoginInfo
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"注销登录"`
}

type LogoutRes struct{}

type RestPasswordReq struct {
	g.Meta `path:"/resetPass" method:"post" summary:"重置密码"`
	Id     int `json:"id"`
}

type RestPasswordRes struct{}
