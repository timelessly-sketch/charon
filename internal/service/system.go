// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon/internal/model"
	"charon/internal/model/entity"
	"context"
)

type (
	IConfig interface {
		InitConfig(ctx context.Context) (err error)
		LoadConfig(ctx context.Context) (err error)
		LoadAuthApiPath(ctx context.Context) (err error)
		LoadAuthMenu(ctx context.Context) (err error)
		LoadToken(ctx context.Context) (cfg *model.Token, err error)
	}
	ISystem interface {
		RoleList(ctx context.Context) (records []entity.Role, err error)
		MenuList(ctx context.Context) (records []entity.Menu, err error)
		MenuDynamic(ctx context.Context, id int) (records []entity.Menu, err error)
		MenuByName(ctx context.Context, name string) (record entity.Menu, err error)
		MenuAdd(ctx context.Context, menu entity.Menu) (err error)
		MenuEdit(ctx context.Context, menu entity.Menu) (err error)
		ApiList(ctx context.Context) (records []entity.Api, err error)
		ApiEdit(ctx context.Context, api entity.Api) (err error)
		ApiAdd(ctx context.Context, api entity.Api) (err error)
		ApiByName(ctx context.Context, name string) (api entity.Api, err error)
	}
)

var (
	localConfig IConfig
	localSystem ISystem
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

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
