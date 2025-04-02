package system

import (
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta   `path:"/userList" method:"get" summary:"用户设置"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	public.Pagination
}

type UserListRes struct {
	Records []entity.User `json:"records"`
	Total   int           `json:"total"`
}

type UserEditReq struct {
	g.Meta `path:"/userEdit" method:"post" summary:"编辑用户"`
	entity.User
}

type UserEditRes struct{}
