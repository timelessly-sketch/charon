package system

import (
	"charon/internal/consts"
	"charon/internal/dao"
	"charon/internal/library/cache"
	"charon/internal/library/token"
	"charon/internal/model/entity"
	"charon/internal/model/public"
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
	if err := s.LoadAuthApiPath(ctx); err != nil {
		return err
	}
	if err := s.LoadAuthMenu(ctx); err != nil {
		return err
	}

	cfg, err := s.LoadToken(ctx)
	if err != nil {
		return err
	}

	token.SetConfig(cfg)
	return
}

func (s *sConfig) LoadAuthApiPath(ctx context.Context) (err error) {
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
			path := consts.BuildPathMethod(api.Path, api.Method)
			if api.Roles == nil {
				permMap[path] = true
				continue
			}
			arrayFrom := garray.NewArrayFrom(gconv.SliceAny(api.Roles), true)
			permMap[path] = arrayFrom.Contains(role.Name)
		}

		key := cache.BuildRole(role.Name)
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
	g.Log().Info(ctx, "load auth path success")
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

		key := cache.BuildMenu(role.Name)
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
	g.Log().Info(ctx, "load auth menu success")
	return err
}

func (s *sConfig) LoadToken(ctx context.Context) (cfg *public.Token, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&cfg)
	return
}
