package cluster

import (
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/clusterList" method:"get" summary:"获取集群列表"`
	public.Pagination
}

type ListRes struct {
	Records []entity.Cluster `json:"records"`
	Total   int              `json:"total"`
}

type EditReq struct {
	g.Meta `path:"/clusterEdit" method:"post" summary:"编辑集群信息"`
	entity.Cluster
}
type EditRes struct{}
