package public

type LoginInput struct {
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

type LoginInfo struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
