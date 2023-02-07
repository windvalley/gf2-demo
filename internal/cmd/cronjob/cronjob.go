package cronjob

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"gf2-demo/utility"
)

var (
	Main = gcmd.Command{
		Name:        "gf2-demo-job",
		Brief:       "",
		Description: "",
		Usage: `Dev:
		./gf2-demo-job

	Test:
		./gf2-demo-job --gf.gcfg.file=config.test.yaml
		or 
		export GF_GCFG_FILE=config.test.yaml && ./gf2-demo-job

	Prod:
		./gf2-demo-job --gf.gcfg.file=config.prod.yaml
		or 
		export GF_GCFG_FILE=config.prod.yaml && ./gf2-demo-job`,
		Arguments: []gcmd.Argument{
			{
				Name:   "version",
				Short:  "v",
				Brief:  "print version info",
				IsArg:  false,
				Orphan: true,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ver := parser.GetOpt("version")
			if ver != nil {
				utility.PrintVersionInfo()
				return
			}

			// json格式日志
			logFormat, err := g.Cfg().Get(ctx, "logger.cronjob.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 显示打印错误的文件行号
			g.Log().SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)

			// 查看使用的配置文件是哪个
			configFile := g.Cfg().GetAdapter()
			g.Log().Debugf(ctx, "use config file: %+v", configFile)

			// ****************** 以下部分为业务逻辑

			fmt.Printf("gf2-demo-job\n")

			g.Log("cronjob").Info(ctx, "foo")

			g.Log("cronjob").Error(ctx, errors.New("bar"))

			return
		},
	}
)
