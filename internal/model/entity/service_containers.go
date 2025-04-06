// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ServiceContainers is the golang structure for table service_containers.
type ServiceContainers struct {
	ServiceId  string `json:"serviceId"  orm:"service_id" description:"外键关联 tb_service_base.id"` // 外键关联 tb_service_base.id
	Namespace  string `json:"namespace"  orm:"namespace"  description:"K8s命名空间"`                 // K8s命名空间
	Containers string `json:"containers" orm:"containers" description:"多容器管理"`                   // 多容器管理
}
