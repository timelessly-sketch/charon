package service

import (
	"charon/internal/consts"
	"charon/internal/dao"
	"charon/internal/model/public"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sService struct{}

func NewService() *sService {
	return &sService{}
}

func init() {
	service.RegisterService(NewService())
}

func (s *sService) List(ctx context.Context, serviceType consts.ServiceType, environment, name string, page public.Pagination) (records any, total int, err error) {
	var (
		containerList []public.KubernetesService
	)

	db := dao.ServiceBase.Ctx(ctx).OmitEmpty().Where(g.Map{
		"environment": environment,
		"NAME LIKE":   "%" + name + "%"})
	total, err = db.Count()
	err = db.Limit((page.Page-1)*page.Size, page.Size).ScanList(&containerList, "ServiceBase")

	switch serviceType {
	case consts.Kubernetes:
		err = dao.ServiceContainers.Ctx(ctx).Where(
			"service_id", gdb.ListItemValuesUnique(containerList, "ServiceBase", "BaseId")).
			ScanList(&containerList, "ServiceContainers", "ServiceBase", "service_id:base_id")
		if err != nil {
			return
		}
		return containerList, total, nil
	case consts.Cvm:

	}
	return
}
