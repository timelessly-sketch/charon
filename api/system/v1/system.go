package v1

import (
	"charon/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta  `path:"/roleList" method:"get" summary:"获取全部角色" role:"admin,operator"`
	Current int `json:"current"`
	Size    int `json:"size"`
}

type RoleListRes struct {
	Records []entity.Role `json:"records"`
}

type UserListReq struct {
	g.Meta   `path:"/userPage" method:"get" summary:"用户设置" role:"admin,operator"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"10"`
	UserName string `json:"username"`
	Name     string `json:"name"`
}

type UserListRes struct {
	Records []entity.User `json:"records"`
	Total   int           `json:"total"`
}
