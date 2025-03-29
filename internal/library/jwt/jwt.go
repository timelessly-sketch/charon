package jwt

import (
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/model"
	"charon/internal/model/entity"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var cfg *model.Token

func SetConfig(c *model.Token) {
	cfg = c
}

type CustomClaims struct {
	entity.User
	jwt.RegisteredClaims
}

func GenerateJWT(ctx context.Context, user entity.User) (token string, err error) {
	var (
		now      = time.Now()
		duration = time.Second * gconv.Duration(cfg.ExpiresAt)
	)

	v, err := cache.Instance().Get(ctx, user.UserName+":Token")
	if err != nil {
		g.Log().Error(ctx, err)
		return "", gerror.NewCode(consts.CodeRedisError)
	}
	if !v.IsEmpty() && !cfg.MultiLogin {
		return "", gerror.NewCode(consts.CodeUserNotMultiLogin)
	}

	claims := CustomClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    cfg.Issuer,
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(cfg.SecretKey)
	if err != nil {
		g.Log().Error(ctx, err)
		return "", gerror.NewCode(consts.CodeGenerateToken)
	}
	_ = cache.Instance().Set(ctx, user.UserName+":Token", token, duration)
	return
}

func ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return cfg.SecretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
