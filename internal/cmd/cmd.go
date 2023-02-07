package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/glog"

	"gf2-api/internal/controller"
	"gf2-api/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// json格式日志
			logFormat, err := g.Cfg().Get(ctx, "logger.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 异步打印日志 & 显示打印错误的文件行号, 对访问日志无效
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_STD | glog.F_FILE_LONG)

			// 简化报错堆栈, 使不包含gf框架的部分, 仅对访问日志有效
			envGfGerrorBrief := genv.Get("GF_GERROR_BRIEF")
			g.Log().Debugf(ctx, "GF_GERROR_BRIEF=%s", envGfGerrorBrief)

			s := g.Server()
			s.Use(
				service.Middleware().ResponseHandler,
			)
			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.Hello,
				)
			})
			s.Run()
			return nil
		},
	}
)
