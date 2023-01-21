package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func (s *sMiddleware) CorsHandler(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
