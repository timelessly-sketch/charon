// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for the table tb_menu.
type MenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuColumns defines and stores column names for the table tb_menu.
type MenuColumns struct {
	Id            string // 菜单ID
	Pid           string // 父菜单ID
	Name          string // 菜单名称（英文标识）
	Path          string // 路由路径（全局唯一）
	Title         string // 显示标题
	RequiresAuth  string // 是否需要鉴权
	Icon          string // 图标类名
	MenuType      string // 菜单类型：目录/页面
	ComponentPath string // 组件文件路径
	PinTab        string // 是否固定标签页
	Hide          string // 是否隐藏菜单
	ActiveMenu    string // 激活显示的菜单路径
	KeepAlive     string // 是否缓存页面
	WithoutTab    string // 当前路由是否会被添加到Tab中
	Href          string // 外部链接地址
	Order         string // 菜单排序权重
	Roles         string // 允许访问的角色列表
	CreatedAt     string // 创建时间
	UpdatedBy     string // 更新者
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// menuColumns holds the columns for the table tb_menu.
var menuColumns = MenuColumns{
	Id:            "id",
	Pid:           "pid",
	Name:          "name",
	Path:          "path",
	Title:         "title",
	RequiresAuth:  "requires_auth",
	Icon:          "icon",
	MenuType:      "menu_type",
	ComponentPath: "component_path",
	PinTab:        "pin_tab",
	Hide:          "hide",
	ActiveMenu:    "active_menu",
	KeepAlive:     "keep_alive",
	WithoutTab:    "without_tab",
	Href:          "href",
	Order:         "order",
	Roles:         "roles",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao(handlers ...gdb.ModelHandler) *MenuDao {
	return &MenuDao{
		group:    "default",
		table:    "tb_menu",
		columns:  menuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
