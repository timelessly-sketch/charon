// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClusterDao is the data access object for the table tb_cluster.
type ClusterDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ClusterColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ClusterColumns defines and stores column names for the table tb_cluster.
type ClusterColumns struct {
	Id          string // 集群唯一ID
	Name        string // 集群名称(唯一)
	ApiServer   string // API Server地址
	AuthConfig  string // 认证配置(使用token认证)
	Environment string // 所属环境
	Region      string // 物理区域/地理位置
	Status      string // 集群状态 1-正常 2-异常
	AutoMated   string // 1 - 自动发布 2 - 手动发布
	Remark      string // 备注信息
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	UpdatedBy   string // 更新人
	DeletedAt   string // 删除时间
}

// clusterColumns holds the columns for the table tb_cluster.
var clusterColumns = ClusterColumns{
	Id:          "id",
	Name:        "name",
	ApiServer:   "api_server",
	AuthConfig:  "auth_config",
	Environment: "environment",
	Region:      "region",
	Status:      "status",
	AutoMated:   "auto_mated",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	UpdatedBy:   "updated_by",
	DeletedAt:   "deleted_at",
}

// NewClusterDao creates and returns a new DAO object for table data access.
func NewClusterDao(handlers ...gdb.ModelHandler) *ClusterDao {
	return &ClusterDao{
		group:    "default",
		table:    "tb_cluster",
		columns:  clusterColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ClusterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ClusterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ClusterDao) Columns() ClusterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ClusterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ClusterDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ClusterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
