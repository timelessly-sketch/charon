package cluster

import (
	"charon/internal/dao"
	"charon/internal/library/chorm"
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
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

// Edit 新增/编辑
func (s *sCluster) Edit(ctx context.Context, cluster entity.Cluster) (err error) {
	if err = s.VerifyUnique(ctx, &public.VerifyUnique{
		Id: cluster.Id,
		Where: g.Map{
			dao.Cluster.Columns().Name:        cluster.Name,
			dao.Cluster.Columns().Environment: cluster.Environment,
		},
	}); err != nil {
		return err
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if cluster.Id > 0 {
			_, err = dao.Cluster.Ctx(ctx).OmitEmpty().Where(dao.Cluster.Columns().Id, cluster.Id).Data(&cluster).Update()
			return
		}

		if _, err = dao.Cluster.Ctx(ctx).Insert(&cluster); err != nil {
			return err
		}
		return
	})
}

// VerifyUnique 验证集群唯一属性
func (s *sCluster) VerifyUnique(ctx context.Context, in *public.VerifyUnique) (err error) {
	if in.Where == nil {
		return
	}
	cols := dao.Cluster.Columns()
	msgMap := g.MapStrStr{
		cols.Name:        "集群名字已存在，请更换",
		cols.Environment: "集群环境已存在，请更换",
	}
	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = chorm.IsUnique(ctx, &dao.Cluster, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}
