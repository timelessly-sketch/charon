// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon/internal/model/entity"
	"context"
)

type (
	ISystem interface {
		RoleList(ctx context.Context) (records []entity.Role, err error)
		MenuList(ctx context.Context) (records []entity.Menu, err error)
		MenuByName(ctx context.Context, name string) (record entity.Menu, err error)
		MenuCreate(ctx context.Context, menu entity.Menu) (err error)
		MenuEdit(ctx context.Context, menu entity.Menu) (err error)
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
