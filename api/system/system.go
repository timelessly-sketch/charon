// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"charon/api/system/v1"
)

type ISystemV1 interface {
	MenuList(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error)
	MenuAdd(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error)
	MenuEdit(ctx context.Context, req *v1.MenuEditReq) (res *v1.MenuEditRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserEdit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserEditRes, err error)
	UserAdd(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error)
}
