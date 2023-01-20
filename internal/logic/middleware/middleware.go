package middleware

import "goframe-erp-v1/internal/service"

type sMiddleware struct {
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(New())
}
