package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
)

const (
	// NOTE: Do Not Edit the Value.
	// All browser auto generate this header.
	// Use g.Client() to request other service will take this header by default.
	clientTraceIDHeader = "X-Request-Id"
)

// TraceID use X-Request-Id from client request
func (s *sMiddleware) TraceID(r *ghttp.Request) {
	traceID := r.GetHeader(clientTraceIDHeader)
	if traceID != "" {
		newCtx, err := gtrace.WithUUID(r.Context(), traceID)
		if err != nil {
			panic(err)
		}

		r.SetCtx(newCtx)
	}

	r.Middleware.Next()
}
