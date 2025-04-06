// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ServiceBase is the golang structure for table service_base.
type ServiceBase struct {
	Id          int         `json:"id"          orm:"id"          description:"自增ID"`          // 自增ID
	BaseId      string      `json:"baseId"      orm:"base_id"     description:"服务唯一ID (UUID)"` // 服务唯一ID (UUID)
	Name        string      `json:"name"        orm:"name"        description:"服务名称"`          // 服务名称
	Type        string      `json:"type"        orm:"type"        description:"服务类型"`          // 服务类型
	Status      string      `json:"status"      orm:"status"      description:"服务状态"`          // 服务状态
	Environment string      `json:"environment" orm:"environment" description:"部署环境"`          // 部署环境
	Project     string      `json:"project"     orm:"project"     description:"所属项目组"`         // 所属项目组
	GitUrl      string      `json:"gitUrl"      orm:"git_url"     description:"代码仓库信息"`        // 代码仓库信息
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`          // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`          // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`          // 删除时间
	Description string      `json:"description" orm:"description" description:"服务描述"`          // 服务描述
}
