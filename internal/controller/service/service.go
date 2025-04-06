package service

import (
	"charon/api/service"
	serviceLogic "charon/internal/service"
	"context"
)

var (
	Service = cService{}
)

type cService struct{}

func (c *cService) List(ctx context.Context, req *service.ListReq) (res *service.ListRes, err error) {
	records, total, err := serviceLogic.Service().List(ctx, req.Type, req.Environment, req.Name, req.Pagination)
	res = &service.ListRes{
		Records: records,
		Total:   total,
	}
	return
}
