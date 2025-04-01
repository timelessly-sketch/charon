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
	ICluster interface {
		List(ctx context.Context, page int, size int) (records []entity.Cluster, total int, err error)
		Edit(ctx context.Context, cluster entity.Cluster) (err error)
	}
)

var (
	localCluster ICluster
)

func Cluster() ICluster {
	if localCluster == nil {
		panic("implement not found for interface ICluster, forgot register?")
	}
	return localCluster
}

func RegisterCluster(i ICluster) {
	localCluster = i
}
