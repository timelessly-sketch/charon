// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         int         `json:"id"         orm:"id"         description:"自增ID"`       // 自增ID
	NickName   string      `json:"nickName"   orm:"nick_name"  description:"昵称"`         // 昵称
	RealName   string      `json:"realName"   orm:"real_name"  description:"用户名"`        // 用户名
	RoleName   string      `json:"roleName"   orm:"role_name"  description:"角色组"`        // 角色组
	Email      string      `json:"email"      orm:"email"      description:"邮箱"`         // 邮箱
	Status     bool        `json:"status"     orm:"status"     description:"0-禁用, 1-正常"` // 0-禁用, 1-正常
	AvatarUrl  string      `json:"avatarUrl"  orm:"avatar_url" description:"头像"`         // 头像
	Department string      `json:"department" orm:"department" description:"部门"`         // 部门
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`       // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`       // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" description:"删除时间"`       // 删除时间
}
