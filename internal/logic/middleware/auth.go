package middleware

import (
	"charon/internal/consts"
	"charon/internal/library/cache"
	"charon/internal/library/token"
	"charon/internal/model/entity"
	"charon/internal/service"
	"context"
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
		handler    = r.GetServeHandler()
		route      = handler.Handler.Router.Uri
		method     = r.Method
		path       = route + ":" + method
		apiNotAuth = g.Map{"code": http.StatusForbidden, "message": "api not auth"}
	)

	if handler.GetMetaTag("noAuth") == "true" || method == "OPTIONS" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized, consts.CodeTokenMissError)
		return
	}
	tokenString := authHeader[len("Bearer "):]
	claims, err := token.ValidateJWT(tokenString)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusUnauthorized, consts.CodeTokenInvalidError)
		return
	}

	key := cache.BuildRole(claims.RoleName)
	value, _ := cache.Instance().Get(r.Context(), key)
	if k, ok := value.Map()[path]; !gconv.Bool(k) || !ok {
		r.Response.WriteStatusExit(http.StatusOK, apiNotAuth)
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
