// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id          int         `json:"id"          orm:"id"          description:"自增id"`          // 自增id
	Name        string      `json:"name"        orm:"name"        description:"角色名称"`          // 角色名称
	Description string      `json:"description" orm:"description" description:"描述"`            // 描述
	Status      bool        `json:"status"      orm:"status"      description:"0 - 禁止 1 - 启用"` // 0 - 禁止 1 - 启用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`          // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`          // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`          // 删除时间
}
