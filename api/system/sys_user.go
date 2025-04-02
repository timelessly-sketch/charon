package system

import (
	"charon/internal/model"
	"charon/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta   `path:"/userList" method:"get" summary:"用户设置"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	model.Pagination
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

type UserAddReq struct {
	g.Meta `path:"/userAdd" method:"post" summary:"添加用户"`
	entity.User
}

type UserAddRes struct{}
