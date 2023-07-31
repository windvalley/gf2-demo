package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/windvalley/gf2-demo/internal/consts"
)

// AccessUser set access user and mail to Context.
// Usage:
//
//	Get user: r.GetCtxVar(consts.CtxAccessUserKey).String()
//	Get user's mail: r.GetCtxVar(consts.CtxAccessUserMailKey).String()
func (s *sMiddleware) AccessUser(r *ghttp.Request) {
	user := r.GetHeader(consts.AccessUserHeader)
	mail := r.GetHeader(consts.AccessUserMailHeader)

	if user != "" {
		r.SetCtxVar(consts.CtxAccessUserKey, user)
	}

	if mail != "" {
		r.SetCtxVar(consts.CtxAccessUserMailKey, mail)
	}

	r.Middleware.Next()
}
