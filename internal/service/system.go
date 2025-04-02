// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"context"
)

type (
	IConfig interface {
		InitConfig(ctx context.Context) (err error)
		LoadConfig(ctx context.Context) (err error)
		LoadAuthApiPath(ctx context.Context) (err error)
		LoadAuthMenu(ctx context.Context) (err error)
		LoadToken(ctx context.Context) (cfg *public.Token, err error)
	}
	ISysApi interface {
		List(ctx context.Context) (records []entity.Api, err error)
		Edit(ctx context.Context, api entity.Api) (err error)
	}
	ISysMenu interface {
		RoleList(ctx context.Context) (records []entity.Role, err error)
		MenuList(ctx context.Context, id int) (records []entity.Menu, err error)
		MenuEdit(ctx context.Context, menu *entity.Menu) (err error)
	}
	ISysUser interface {
		// List 获取用户列表
		List(ctx context.Context, username string, name string, page int, size int) (records []entity.User, total int, err error)
		// Edit 编辑或新增用户
		Edit(ctx context.Context, user *entity.User) (err error)
		ResetPassword(ctx context.Context, id int) (err error)
		Login(ctx context.Context, input public.LoginInput) (info public.LoginInfo, err error)
		// VerifyUnique 验证用户唯一属性
		VerifyUnique(ctx context.Context, in *public.VerifyUnique) (err error)
	}
)

var (
	localConfig  IConfig
	localSysApi  ISysApi
	localSysMenu ISysMenu
	localSysUser ISysUser
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}

func SysApi() ISysApi {
	if localSysApi == nil {
		panic("implement not found for interface ISysApi, forgot register?")
	}
	return localSysApi
}

func RegisterSysApi(i ISysApi) {
	localSysApi = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
