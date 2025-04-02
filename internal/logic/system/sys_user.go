package system

import (
	"charon/internal/consts"
	"charon/internal/dao"
	"charon/internal/library/chorm"
	"charon/internal/library/token"
	"charon/internal/model/entity"
	"charon/internal/model/public"
	"charon/internal/service"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysUser struct{}

func NewSysUser() *sSysUser {
	return &sSysUser{}
}

func init() {
	service.RegisterSysUser(NewSysUser())
}

// List 获取用户列表
func (s *sSysUser) List(ctx context.Context, username, name string, page, size int) (records []entity.User, total int, err error) {
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

// Edit 编辑或新增用户
func (s *sSysUser) Edit(ctx context.Context, user *entity.User) (err error) {
	if err = s.VerifyUnique(ctx, &public.VerifyUnique{
		Id: user.Id,
		Where: g.Map{
			dao.User.Columns().NickName: user.NickName,
			dao.User.Columns().UserName: user.UserName,
			dao.User.Columns().UserId:   user.UserId,
			dao.User.Columns().Email:    user.Email,
			dao.User.Columns().Phone:    user.Phone,
		},
	}); err != nil {
		return err
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if user.Id > 0 {
			_, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(&user).Update()
			return
		}
		user.Password = s.generatePassword(user.UserName)
		if _, err = dao.User.Ctx(ctx).Insert(&user); err != nil {
			return err
		}
		return
	})
}

func (s *sSysUser) ResetPassword(ctx context.Context, id int) (err error) {
	var user entity.User
	if err = dao.User.Ctx(ctx).WherePri(id).Scan(&user); err != nil || gerror.Is(err, sql.ErrNoRows) {
		return gerror.NewCode(consts.CodeUserNotFound)
	}
	user.Password = s.generatePassword(user.NickName)
	if _, err = dao.User.Ctx(ctx).WherePri(id).Data(&user).Update(); err != nil {
		return gerror.NewCode(consts.CodeResetPassError)
	}
	return
}

func (s *sSysUser) Login(ctx context.Context, input public.LoginInput) (info public.LoginInfo, err error) {
	var (
		user = entity.User{}
	)

	if err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, input.Username).Scan(&user); err != nil || gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return info, gerror.NewCode(consts.CodeUserNotFound)
	}
	if s.generatePassword(input.Password) != user.Password {
		return info, gerror.NewCode(consts.CodePasswordInvalid)
	}
	generateJWT, err := token.GenerateJWT(ctx, user)
	if err != nil {
		g.Log().Warning(ctx, err)
		return
	}

	info = public.LoginInfo{
		Id:       user.Id,
		Username: user.UserName,
		Avatar:   user.AvatarUrl,
		Token:    generateJWT,
		Role:     user.RoleName,
	}

	return
}

// VerifyUnique 验证用户唯一属性
func (s *sSysUser) VerifyUnique(ctx context.Context, in *public.VerifyUnique) (err error) {
	if in.Where == nil {
		return
	}
	cols := dao.User.Columns()
	msgMap := g.MapStrStr{
		cols.NickName: "昵称已存在，请更换",
		cols.UserName: "英文名已存在，请更换",
		cols.UserId:   "用户ID已存在，请更换",
		cols.Email:    "用户已存在，请更换",
		cols.Phone:    "phone已存在，请跟换",
	}
	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = chorm.IsUnique(ctx, &dao.Cluster, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

func (s *sSysUser) generatePassword(password string) string {
	hash := sha256.New()
	hash.Write(gconv.Bytes(password))
	return hex.EncodeToString(hash.Sum(nil))
}
