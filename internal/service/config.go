// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"charon/internal/model"
	"context"
)

type (
	IConfig interface {
		InitConfig(ctx context.Context) (err error)
		LoadConfig(ctx context.Context) (err error)
		LoadToken(ctx context.Context) (cfg *model.Token, err error)
	}
)

var (
	localConfig IConfig
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
