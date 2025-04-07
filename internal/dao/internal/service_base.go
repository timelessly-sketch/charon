// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServiceBaseDao is the data access object for the table tb_service_base.
type ServiceBaseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ServiceBaseColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ServiceBaseColumns defines and stores column names for the table tb_service_base.
type ServiceBaseColumns struct {
	Id          string // 自增ID
	BaseId      string // 服务唯一ID (UUID)
	Name        string // 服务名称
	ServiceType string // 服务类型
	Status      string // 服务状态
	Environment string // 部署环境
	Project     string // 所属项目组
	GitUrl      string // 代码仓库信息
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	Description string // 服务描述
}

// serviceBaseColumns holds the columns for the table tb_service_base.
var serviceBaseColumns = ServiceBaseColumns{
	Id:          "id",
	BaseId:      "base_id",
	Name:        "name",
	ServiceType: "service_type",
	Status:      "status",
	Environment: "environment",
	Project:     "project",
	GitUrl:      "git_url",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Description: "description",
}

// NewServiceBaseDao creates and returns a new DAO object for table data access.
func NewServiceBaseDao(handlers ...gdb.ModelHandler) *ServiceBaseDao {
	return &ServiceBaseDao{
		group:    "default",
		table:    "tb_service_base",
		columns:  serviceBaseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ServiceBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ServiceBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ServiceBaseDao) Columns() ServiceBaseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ServiceBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ServiceBaseDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ServiceBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
