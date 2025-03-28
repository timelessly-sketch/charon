// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table tb_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta        `orm:"table:tb_menu, do:true"`
	Id            interface{} // 菜单ID
	Pid           interface{} // 父菜单ID
	Name          interface{} // 菜单名称（英文标识）
	Path          interface{} // 路由路径（全局唯一）
	Title         interface{} // 显示标题
	RequiresAuth  interface{} // 是否需要鉴权
	Icon          interface{} // 图标类名
	MenuType      interface{} // 菜单类型：目录/页面
	ComponentPath interface{} // 组件文件路径
	PinTab        interface{} // 是否固定标签页
	Hide          interface{} // 是否隐藏菜单
	ActiveMenu    interface{} // 激活显示的菜单路径
	KeepAlive     interface{} // 是否缓存页面
	WithoutTab    interface{} // 当前路由是否会被添加到Tab中
	Href          interface{} // 外部链接地址
	Order         interface{} // 菜单排序权重
	Roles         *gjson.Json // 允许访问的角色列表
	CreatedAt     *gtime.Time // 创建时间
	UpdatedBy     interface{} // 更新者
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
