package system

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
)

type sSystem struct{}

func NewSystem() *sSystem {
	return &sSystem{}
}

func init() {
	service.RegisterSystem(NewSystem())
}

func (sys *sSystem) RoleList(ctx context.Context) (records []entity.Role, err error) {
	err = dao.Role.Ctx(ctx).Scan(&records)
	return
}
