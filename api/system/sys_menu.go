package system

import (
	"charon/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuListReq struct {
	g.Meta `path:"/menuList" method:"get" summary:"获取菜单列表"`
	Id     int `json:"Id" d:"0"`
}

type MenuListRes struct {
	Records []entity.Menu `json:"records"`
}

type MenuEditReq struct {
	g.Meta `path:"/menuEdit" method:"post" summary:"新增菜单"`
	entity.Menu
}

type MenuEditRes struct{}
