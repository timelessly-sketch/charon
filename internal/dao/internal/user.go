// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table tb_user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table tb_user.
type UserColumns struct {
	Id         string // 自增ID
	NickName   string // 昵称
	UserName   string // 英文名
	Name       string // 中文名
	Department string // 部门
	UserId     string // 用户Id
	RoleName   string // 角色组
	Email      string // 邮箱
	Phone      string // 电话
	Status     string // 0-禁用, 1-正常
	AvatarUrl  string // 头像
	CreatedAt  string // 创建时间
	UpdatedBy  string // 更新人
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// userColumns holds the columns for the table tb_user.
var userColumns = UserColumns{
	Id:         "id",
	NickName:   "nick_name",
	UserName:   "user_name",
	Name:       "name",
	Department: "department",
	UserId:     "user_id",
	RoleName:   "role_name",
	Email:      "email",
	Phone:      "phone",
	Status:     "status",
	AvatarUrl:  "avatar_url",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "tb_user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
