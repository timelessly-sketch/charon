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
	IUser interface {
		Select(ctx context.Context, data entity.User) (record entity.User, err error)
		List(ctx context.Context, username string, name string, page int, size int) (records []entity.User, total int, err error)
		Edit(ctx context.Context, user entity.User) (err error)
		Create(ctx context.Context, user entity.User) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
