package middleware

import (
	"charon/internal/consts"
	"charon/internal/library/jwt"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

type sMiddleware struct{}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func (m *sMiddleware) AuthMiddleware(r *ghttp.Request) {
	handler := r.GetServeHandler()
	noAuth := handler.GetMetaTag("noAuth")
	if noAuth == "true" || r.Request.Method == "OPTIONS" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(http.StatusInternalServerError, consts.CodeTokenInvalid)
		return
	}
	tokenString := authHeader[len("Bearer "):]
	claims, err := jwt.ValidateJWT(tokenString)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, consts.CodeTokenInvalid)
		return
	}
	r.SetCtxVar("user", claims.User)
	auth := handler.GetMetaTag("role")
	fmt.Println(auth)
	if !gstr.Contains(auth, claims.RoleName) {
		r.Response.WriteStatusExit(http.StatusUnauthorized, gcode.CodeInvalidOperation)
		return
	}
	r.Middleware.Next()
}

// CORS 跨域
func (m *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (m *sMiddleware) GetCtxUser(ctx context.Context) (user *entity.User) {
	if err := g.RequestFromCtx(ctx).GetCtxVar("user").Scan(&user); err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	return
}
