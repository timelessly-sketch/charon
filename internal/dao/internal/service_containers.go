// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServiceContainersDao is the data access object for the table tb_service_containers.
type ServiceContainersDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  ServiceContainersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// ServiceContainersColumns defines and stores column names for the table tb_service_containers.
type ServiceContainersColumns struct {
	ServiceId  string // 外键关联 tb_service_base.id
	Namespace  string // K8s命名空间
	Containers string // 多容器管理
}

// serviceContainersColumns holds the columns for the table tb_service_containers.
var serviceContainersColumns = ServiceContainersColumns{
	ServiceId:  "service_id",
	Namespace:  "namespace",
	Containers: "containers",
}

// NewServiceContainersDao creates and returns a new DAO object for table data access.
func NewServiceContainersDao(handlers ...gdb.ModelHandler) *ServiceContainersDao {
	return &ServiceContainersDao{
		group:    "default",
		table:    "tb_service_containers",
		columns:  serviceContainersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ServiceContainersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ServiceContainersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ServiceContainersDao) Columns() ServiceContainersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ServiceContainersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ServiceContainersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ServiceContainersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
