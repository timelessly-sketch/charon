// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApiDao is the data access object for the table tb_api.
type ApiDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ApiColumns         // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ApiColumns defines and stores column names for the table tb_api.
type ApiColumns struct {
	Id        string // 接口ID
	Pid       string // 接口PID
	Name      string // 接口名
	Icon      string // 图标
	Title     string // 标题
	Path      string // 接口路径
	Method    string // 接口方法-目录为空,接口不能为空
	ApiType   string // 接口或者目录
	Roles     string // 权限列表
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// apiColumns holds the columns for the table tb_api.
var apiColumns = ApiColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Icon:      "icon",
	Title:     "title",
	Path:      "path",
	Method:    "method",
	ApiType:   "api_type",
	Roles:     "roles",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewApiDao creates and returns a new DAO object for table data access.
func NewApiDao(handlers ...gdb.ModelHandler) *ApiDao {
	return &ApiDao{
		group:    "default",
		table:    "tb_api",
		columns:  apiColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ApiDao) Columns() ApiColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ApiDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
