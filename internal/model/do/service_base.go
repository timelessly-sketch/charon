// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ServiceBase is the golang structure of table tb_service_base for DAO operations like Where/Data.
type ServiceBase struct {
	g.Meta      `orm:"table:tb_service_base, do:true"`
	Id          interface{} // 自增ID
	BaseId      interface{} // 服务唯一ID (UUID)
	Name        interface{} // 服务名称
	ServiceType interface{} // 服务类型
	Status      interface{} // 服务状态
	Environment interface{} // 部署环境
	Project     interface{} // 所属项目组
	GitUrl      interface{} // 代码仓库信息
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Description interface{} // 服务描述
}
