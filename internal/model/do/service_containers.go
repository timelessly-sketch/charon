// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ServiceContainers is the golang structure of table tb_service_containers for DAO operations like Where/Data.
type ServiceContainers struct {
	g.Meta     `orm:"table:tb_service_containers, do:true"`
	ServiceId  interface{} // 外键关联 tb_service_base.id
	Namespace  interface{} // K8s命名空间
	Containers interface{} // 多容器管理
}
