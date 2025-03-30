// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// K8SClusterDao is the data access object for the table tb_k8s_cluster.
type K8SClusterDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  K8SClusterColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// K8SClusterColumns defines and stores column names for the table tb_k8s_cluster.
type K8SClusterColumns struct {
	Id          string // 集群唯一ID
	Name        string // 集群名称(唯一)
	ApiServer   string // API Server地址
	AuthConfig  string // 认证配置(使用token认证)
	Environment string // 所属环境
	Region      string // 物理区域/地理位置
	Status      string // 集群状态 1-正常 2-异常
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	UpdatedBy   string // 更新人
	DeletedAt   string // 删除时间
	Remark      string // 备注信息
}

// k8SClusterColumns holds the columns for the table tb_k8s_cluster.
var k8SClusterColumns = K8SClusterColumns{
	Id:          "id",
	Name:        "name",
	ApiServer:   "api_server",
	AuthConfig:  "auth_config",
	Environment: "environment",
	Region:      "region",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	UpdatedBy:   "updated_by",
	DeletedAt:   "deleted_at",
	Remark:      "remark",
}

// NewK8SClusterDao creates and returns a new DAO object for table data access.
func NewK8SClusterDao(handlers ...gdb.ModelHandler) *K8SClusterDao {
	return &K8SClusterDao{
		group:    "default",
		table:    "tb_k8s_cluster",
		columns:  k8SClusterColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *K8SClusterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *K8SClusterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *K8SClusterDao) Columns() K8SClusterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *K8SClusterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *K8SClusterDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *K8SClusterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
