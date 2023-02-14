package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
)

const (
	clientTraceIDHeader = "Trace-Id"
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
