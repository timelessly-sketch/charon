package cluster

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

func init() {
	service.RegisterCluster(NewCluster())
}

func (s *sCluster) List(ctx context.Context, page, size int) (records []entity.Cluster, total int, err error) {
	err = dao.Cluster.Ctx(ctx).Limit((page-1)*size, size).ScanAndCount(&records, &total, false)
	return
}

func (s *sCluster) Edit(ctx context.Context, cluster entity.Cluster) (err error) {
	count, err := dao.Cluster.Ctx(ctx).Where(g.Map{
		dao.Cluster.Columns().Id:          cluster.Id,
		dao.Cluster.Columns().Environment: cluster.Environment,
	}).Count()
	if err != nil {
		return err
	}

	if count == 0 && cluster.Id == 0 {
		if _, err = dao.Cluster.Ctx(ctx).Insert(&cluster); err != nil {
			return err
		}
		return
	}

	if _, err = dao.Cluster.Ctx(ctx).OmitEmpty().Where(dao.Cluster.Columns().Id, cluster.Id).Data(&cluster).Update(); err != nil {
		return err
	}
	return
}
