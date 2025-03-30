// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure for table cluster.
type Cluster struct {
	Id          int         `json:"id"          orm:"id"          description:"集群唯一ID"`            // 集群唯一ID
	Name        string      `json:"name"        orm:"name"        description:"集群名称(唯一)"`          // 集群名称(唯一)
	ApiServer   string      `json:"apiServer"   orm:"api_server"  description:"API Server地址"`      // API Server地址
	AuthConfig  string      `json:"authConfig"  orm:"auth_config" description:"认证配置(使用token认证)"`   // 认证配置(使用token认证)
	Environment string      `json:"environment" orm:"environment" description:"所属环境"`              // 所属环境
	Region      string      `json:"region"      orm:"region"      description:"物理区域/地理位置"`         // 物理区域/地理位置
	Status      string      `json:"status"      orm:"status"      description:"集群状态 1-正常 2-异常"`    // 集群状态 1-正常 2-异常
	AutoMated   string      `json:"autoMated"   orm:"auto_mated"  description:"1 - 自动发布 2 - 手动发布"` // 1 - 自动发布 2 - 手动发布
	Remark      string      `json:"remark"      orm:"remark"      description:"备注信息"`              // 备注信息
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`              // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`              // 更新时间
	UpdatedBy   string      `json:"updatedBy"   orm:"updated_by"  description:"更新人"`               // 更新人
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`              // 删除时间
}
