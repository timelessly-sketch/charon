// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure for table api.
type Api struct {
	Id        int         `json:"id"        orm:"id"         description:"接口ID"`             // 接口ID
	Pid       int         `json:"pid"       orm:"pid"        description:"接口PID"`            // 接口PID
	Name      string      `json:"name"      orm:"name"       description:"接口名"`              // 接口名
	Icon      string      `json:"icon"      orm:"icon"       description:"图标"`               // 图标
	Title     string      `json:"title"     orm:"title"      description:"标题"`               // 标题
	Path      string      `json:"path"      orm:"path"       description:"接口路径"`             // 接口路径
	Method    string      `json:"method"    orm:"method"     description:"接口方法-目录为空,接口不能为空"` // 接口方法-目录为空,接口不能为空
	ApiType   string      `json:"apiType"   orm:"api_type"   description:"接口或者目录"`           // 接口或者目录
	Roles     *gjson.Json `json:"roles"     orm:"roles"      description:"权限列表"`             // 权限列表
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`             // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`             // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`             // 删除时间
}
