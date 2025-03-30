// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure of table tb_cluster for DAO operations like Where/Data.
type Cluster struct {
	g.Meta      `orm:"table:tb_cluster, do:true"`
	Id          interface{} // 集群唯一ID
	Name        interface{} // 集群名称(唯一)
	ApiServer   interface{} // API Server地址
	AuthConfig  interface{} // 认证配置(使用token认证)
	Environment interface{} // 所属环境
	Region      interface{} // 物理区域/地理位置
	Status      interface{} // 集群状态 1-正常 2-异常
	AutoMated   interface{} // 1 - 自动发布 2 - 手动发布
	Remark      interface{} // 备注信息
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	UpdatedBy   interface{} // 更新人
	DeletedAt   *gtime.Time // 删除时间
}
