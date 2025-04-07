package service

import (
	"charon/internal/consts"
	"charon/internal/dao"
	"charon/internal/model/public"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
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

func (s *sService) EditKubernetesService(ctx context.Context, containers public.KubernetesService) (err error) {
	if err = s.VerifyUnique(ctx, &public.VerifyUnique{
		Id: containers.ServiceBase.Id,
		Where: g.Map{
			dao.ServiceBase.Columns().Name:        containers.ServiceBase.Name,
			dao.ServiceBase.Columns().Environment: containers.ServiceBase.Environment,
		},
	}); err != nil {
		return err
	}
	if containers.ServiceBase.Id > 0 {
		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			if _, err = dao.ServiceBase.Ctx(ctx).WherePri(containers.ServiceBase.Id).Data(&containers.ServiceBase).Update(); err != nil {
				return err
			}
			if _, err = dao.ServiceContainers.Ctx(ctx).WherePri(containers.ServiceContainers.ServiceId).
				Data(&containers.ServiceContainers).Update(); err != nil {
				return err
			}
			return
		})
	}

	id, _ := uuid.NewRandom()
	containers.ServiceBase.BaseId, containers.ServiceContainers.ServiceId = id.String(), id.String()
	if containers.ServiceContainers.Containers == "" {
		containers.ServiceContainers.Containers = containers.ServiceBase.Name
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.ServiceBase.Ctx(ctx).Data(&containers.ServiceBase).Insert(); err != nil {
			return err
		}
		if _, err = dao.ServiceContainers.Ctx(ctx).Data(&containers.ServiceContainers).Insert(); err != nil {
			return err
		}
		return
	})
}

func (s *sService) VerifyUnique(ctx context.Context, in *public.VerifyUnique) (err error) {
	count, err := dao.ServiceBase.Ctx(ctx).Where(in.Where).WhereNot(dao.ServiceBase.Columns().Id, in.Id).Count()
	if err != nil || count > 0 {
		g.Log().Warning(ctx, err)
		return gerror.NewCode(consts.CodeServiceExists)
	}
	return
}
