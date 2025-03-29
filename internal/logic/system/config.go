package system

import (
	"charon/internal/library/cache"
	"charon/internal/library/jwt"
	"charon/internal/model"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
	// token config
	cfg, err := s.LoadToken(ctx)
	if err != nil {
		return err
	}
	jwt.SetConfig(cfg)
	return
}

func (s *sConfig) LoadToken(ctx context.Context) (cfg *model.Token, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&cfg)
	return
}
