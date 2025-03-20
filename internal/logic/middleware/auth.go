package middleware

import (
	"charon/internal/consts"
	"charon/internal/library/jwt"
	"charon/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
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
	fmt.Println(authHeader, "1111", r.RequestURI)
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
	r.SetCtxVar("name", claims.Name)
	r.Middleware.Next()
}

// CORS 跨域
func (m *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
