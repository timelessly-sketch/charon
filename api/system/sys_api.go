package system

import (
	"charon/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type ApiListReq struct {
	g.Meta `path:"/apiList" method:"get" summary:"获取所有的api接口"`
}
type ApiListRes struct {
	Records []entity.Api `json:"records"`
}

type ApiEditReq struct {
	g.Meta `path:"/apiEdit" method:"post" summary:"编辑api接口"`
	entity.Api
}
type ApiEditRes struct{}

type ApiAddReq struct {
	g.Meta `path:"/apiAdd" method:"post" summary:"新增api接口"`
	entity.Api
}
type ApiAddRes struct{}
