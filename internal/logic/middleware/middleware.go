package middleware

import "github.com/windvalley/gf2-demo/internal/service"

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}
