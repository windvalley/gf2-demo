package middleware

import "gf2-demo/internal/service"

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(new())
}

func new() *sMiddleware {
	return &sMiddleware{}
}
