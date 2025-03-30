package cluster

import (
	"charon/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/clusterList" method:"get" summary:"获取集群列表"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"10"`
}

type ListRes struct {
	Records []entity.Cluster `json:"records"`
	Total   int              `json:"total"`
}
