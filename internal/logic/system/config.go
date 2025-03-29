package system

import (
	"charon/internal/dao"
	"charon/internal/library/cache"
	"charon/internal/library/jwt"
	"charon/internal/model"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sConfig struct{}

func NewConfig() *sConfig {
	return &sConfig{}
}

func init() {
	service.RegisterConfig(NewConfig())
}

func (s *sConfig) InitConfig(ctx context.Context) (err error) {
	cache.SetAdapter(ctx)
	return s.LoadConfig(ctx)
}

func (s *sConfig) LoadConfig(ctx context.Context) (err error) {
	if err := s.LoadAuthPath(ctx); err != nil {
		return err
	}
	if err := s.LoadAuthMenu(ctx); err != nil {
		return err
	}

	cfg, err := s.LoadToken(ctx)
	if err != nil {
		return err
	}

	jwt.SetConfig(cfg)
	return
}

func (s *sConfig) LoadAuthPath(ctx context.Context) (err error) {
	var (
		roleList = make([]*entity.Role, 0)
		apiList  = make([]entity.Api, 0)
	)

	if err := dao.Role.Ctx(ctx).Scan(&roleList); err != nil {
		return err
	}

	if err := dao.Api.Ctx(ctx).Where(dao.Api.Columns().ApiType, "api").Scan(&apiList); err != nil {
		return err
	}

	for _, role := range roleList {
		permMap := make(map[string]bool)
		for _, api := range apiList {
			path := api.Path + ":" + api.Method
			if api.Roles == nil {
				permMap[path] = true
				continue
			}
			arrayFrom := garray.NewArrayFrom(gconv.SliceAny(api.Roles), true)
			permMap[path] = arrayFrom.Contains(role.Name)
		}

		key := role.Name + ":Role"
		value, _ := cache.Instance().Get(ctx, key)
		if value.IsEmpty() {
			if err := cache.Instance().Set(ctx, key, permMap, 0); err != nil {
				return err
			}
			continue
		}
		if _, _, err := cache.Instance().Update(ctx, key, permMap); err != nil {
			return err
		}
	}
	return
}

func (s *sConfig) LoadAuthMenu(ctx context.Context) (err error) {
	var (
		roleList = make([]*entity.Role, 0)
		menuList = make([]entity.Menu, 0)
	)

	if err := dao.Role.Ctx(ctx).Scan(&roleList); err != nil {
		return err
	}
	if err := dao.Menu.Ctx(ctx).Scan(&menuList); err != nil {
		return err
	}

	for _, role := range roleList {
		records := make([]entity.Menu, 0)
		for _, menu := range menuList {
			if menu.Roles == nil {
				records = append(records, menu)
				continue
			}
			arrayFrom := garray.NewArrayFrom(gconv.SliceAny(menu.Roles), true)
			if arrayFrom.Contains(role.Name) {
				records = append(records, menu)
			}
		}
		key := role.Name + ":Menu"
		value, _ := cache.Instance().Get(ctx, key)
		if value.IsEmpty() {
			if err := cache.Instance().Set(ctx, key, records, 0); err != nil {
				return err
			}
			continue
		}
		if _, _, err := cache.Instance().Update(ctx, key, records); err != nil {
			return err
		}
	}

	return err
}

func (s *sConfig) LoadToken(ctx context.Context) (cfg *model.Token, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&cfg)
	return
}
