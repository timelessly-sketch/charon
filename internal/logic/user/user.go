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

func (u *sUser) Select(ctx context.Context, data entity.User) (record entity.User, err error) {
	err = dao.User.Ctx(ctx).OmitEmpty().Where(data).Scan(&record)
	return
}

func (u *sUser) List(ctx context.Context, username, name string, page, size int) (records []entity.User, total int, err error) {
	records = make([]entity.User, 0)
	db := dao.User.Ctx(ctx)
	if username != "" {
		db = db.WhereLike(dao.User.Columns().UserName, "%"+username+"%")
	}
	if name != "" {
		db = db.WhereLike(dao.User.Columns().Name, "%"+name+"%")
	}
	if err := db.Limit((page-1)*size, size).ScanAndCount(&records, &total, false); err != nil {
		return nil, 0, err
	}
	return
}

func (u *sUser) Edit(ctx context.Context, user entity.User) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, user.Id).Data(user).Update()
	return
}

func (u *sUser) Create(ctx context.Context, user entity.User) (err error) {
	_, err = dao.User.Ctx(ctx).OmitEmpty().Insert(user)
	return
}
