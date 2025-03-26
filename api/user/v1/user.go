package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summary:"用户登录" noAuth:"true"`
	Username string `v:"required#请输入用户名" dc:"用户名"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

type LoginRes struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
