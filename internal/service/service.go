// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon/internal/consts"
	"charon/internal/model/public"
	"context"
)

type (
	IService interface {
		List(ctx context.Context, serviceType consts.ServiceType, environment string, name string, page public.Pagination) (records any, total int, err error)
	}
)

var (
	localService IService
)

func Service() IService {
	if localService == nil {
		panic("implement not found for interface IService, forgot register?")
	}
	return localService
}

func RegisterService(i IService) {
	localService = i
}
