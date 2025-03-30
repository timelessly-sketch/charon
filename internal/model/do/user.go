// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table tb_user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:tb_user, do:true"`
	Id         interface{} // 自增ID
	NickName   interface{} // 昵称
	UserName   interface{} // 英文名
	Password   interface{} // 密码
	Name       interface{} // 中文名
	Department interface{} // 部门
	UserId     interface{} // 用户Id
	RoleName   interface{} // 角色组
	Email      interface{} // 邮箱
	Phone      interface{} // 电话
	Status     interface{} // 0-禁用, 1-正常
	AvatarUrl  interface{} // 头像
	Remark     interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedBy  interface{} // 更新人
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
