package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gvalid"

	"gf2-demo/internal/codes"
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	TraceID string      `json:"traceid"`
	Data    interface{} `json:"data"`
}

// ResponseHandler custom response format
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// For /swagger
	if r.Request.URL.Path == "/api.json" {
		return
	}

	// For PProf
	if g.Cfg().MustGet(r.Context(), "server.pprofEnabled").Bool() {
		rPath := g.Cfg().MustGet(r.Context(), "server.pprofPattern").String()
		if rPath == "" {
			rPath = "/debug/pprof"
		}
		if strings.HasPrefix(r.Request.URL.Path, rPath) {
			return
		}
	}

	// Clean exist response info.
	// i.e.:
	//	 1) gf exception recovered info
	//	 2) gf 404 Not Found content
	r.Response.ClearBuffer()

	var (
		err     = r.GetError()
		res     = r.GetHandlerResponse()
		msg     string
		bizCode codes.BizCode
		ok      bool
	)

	if err != nil {
		if _, ok = err.(gvalid.Error); ok { // validation error
			err = gerror.WrapCode(codes.CodeValidationFailed, err)
			code1 := gerror.Code(err)
			bizCode, _ = code1.(codes.BizCode)
		} else {
			code := gerror.Code(err)
			bizCode, ok = code.(codes.BizCode) // custom error codes
			if !ok {
				err = gerror.NewCode(codes.CodeInternal) // internal error
				code1 := gerror.Code(err)
				bizCode, _ = code1.(codes.BizCode)
			}
		}

		msg = err.Error()
	} else {
		if r.Response.Status == 404 { // gf internal 404 error
			bizCode = codes.CodeNotFound.(codes.BizCode)
		} else {
			bizCode = codes.CodeOK.(codes.BizCode)
		}

		msg = bizCode.Message()
	}

	r.Response.WriteHeader(bizCode.BizDetail().HttpCode)
	r.Response.WriteJsonExit(response{
		Code:    bizCode.BizDetail().Code,
		Message: msg,
		TraceID: gctx.CtxId(r.Context()),
		Data:    res,
	})
}
