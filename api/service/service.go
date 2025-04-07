package service

import (
	"charon/internal/consts"
	"charon/internal/model/public"
	"github.com/gogf/gf/v2/frame/g"
)

type EditReq struct {
	g.Meta                   `path:"/edit" method:"post" summary:"编辑服务"`
	ServiceType              consts.ServiceType `json:"serviceType" v:"in:kubernetes,cvm"`
	public.KubernetesService `json:"kubernetesService" v:"required-if:type,kubernetes"`
}

type EditRes struct{}

type ListReq struct {
	g.Meta `path:"/list" method:"get" summary:"获取服务列表"`
	public.Pagination
	ServiceType consts.ServiceType `json:"serviceType" v:"enums" summary:"服务类型"`
	Name        string             `json:"name"`
	Environment string             `json:"environment"`
}

type ListRes struct {
	Records any `json:"records"`
	Total   int `json:"total"`
}
