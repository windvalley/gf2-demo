package middleware

import "gf2-api/internal/service"

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(new())
}

func new() *sMiddleware {
	return &sMiddleware{}
}
