package cluster

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
)

type sCluster struct{}

func NewCluster() *sCluster {
	return &sCluster{}
}

func init() {
	service.RegisterCluster(NewCluster())
}

func (s *sCluster) List(ctx context.Context, page, size int) (records []entity.Cluster, total int, err error) {
	err = dao.Cluster.Ctx(ctx).Scan(&records)
	return
}
