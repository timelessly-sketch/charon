package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summary:"用户登录" noAuth:"true"`
	Username string `v:"required#请输入用户名" dc:"用户名"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type InfoReq struct {
	g.Meta `path:"/getUserInfo" method:"get" summary:"获取用户详情"`
}

type InfoRes struct {
	Id       int      `json:"id"`
	NickName string   `json:"nick_name"`
	Roles    []string `json:"roles"`
}
