package service

import (
	"charon/api/service"
	"charon/internal/consts"
	serviceLogic "charon/internal/service"
	"context"
	"fmt"
)

var (
	Service = cService{}
)

type cService struct{}

func (c *cService) List(ctx context.Context, req *service.ListReq) (res *service.ListRes, err error) {
	records, total, err := serviceLogic.Service().List(ctx, req.ServiceType, req.Environment, req.Name, req.Pagination)
	res = &service.ListRes{
		Records: records,
		Total:   total,
	}
	return
}

func (c *cService) Edit(ctx context.Context, req *service.EditReq) (res *service.EditRes, err error) {
	fmt.Println(req.ServiceType, "类型")
	fmt.Println(req.KubernetesService)
	switch req.ServiceType {
	case consts.Kubernetes:
		err = serviceLogic.Service().EditKubernetesService(ctx, req.KubernetesService)
	}
	return
}
