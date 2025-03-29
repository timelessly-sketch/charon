package middleware

import (
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/library/jwt"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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
	var (
		handler = r.GetServeHandler()
		route   = handler.Handler.Router.Uri
		method  = r.Method
		path    = route + ":" + method
	)
	if handler.GetMetaTag("noAuth") == "true" || method == "OPTIONS" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized, consts.CodeTokenMiss)
		return
	}
	tokenString := authHeader[len("Bearer "):]
	claims, err := jwt.ValidateJWT(tokenString)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, consts.CodeTokenInvalid)
		return
	}

	value, _ := cache.Instance().Get(r.Context(), claims.RoleName+":Role")
	if k, ok := value.Map()[path]; !gconv.Bool(k) || !ok {
		r.Response.WriteStatusExit(http.StatusUnauthorized, gcode.CodeNotAuthorized)
		return
	}

	r.SetCtxVar("user", claims.User)
	r.Middleware.Next()
}

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
