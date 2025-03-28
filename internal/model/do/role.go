// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table tb_role for DAO operations like Where/Data.
type Role struct {
	g.Meta      `orm:"table:tb_role, do:true"`
	Id          interface{} // 自增id
	Name        interface{} // 角色名称
	Description interface{} // 描述
	Remark      interface{} // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   interface{} // 更新者
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
