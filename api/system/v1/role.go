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
