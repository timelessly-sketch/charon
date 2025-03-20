package user

import (
	"charon/internal/dao"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(NewUser())
}

func (u *sUser) UserByNickName(ctx context.Context, nick string) (record entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().NickName, nick).Scan(&record)
	return
}
