package cronjob

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"gf2-demo/utility"
)

var (
	Main = gcmd.Command{
		Name:        "gf2-demo-job",
		Brief:       "",
		Description: "",
		Usage:       "gf2-demo-job [OPTION]",
		Examples: `
			Dev:
				./gf2-demo-job

			Test:
				./gf2-demo-job -c config.test.yaml
				or 
				GF_GCFG_FILE=config.test.yaml ./gf2-demo-job

			Prod:
				./gf2-demo-job -c config.prod.yaml
				or 
				GF_GCFG_FILE=config.prod.yaml ./gf2-demo-job`,
		Additional: "Find more information at: https://github.com/windvalley/gf2-demo",
		Arguments: []gcmd.Argument{
			{
				Name:   "version",
				Short:  "v",
				Brief:  "print version info",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ver := parser.GetOpt("version")
			if ver != nil {
				utility.PrintVersionInfo()
				return
			}

			config := parser.GetOpt("config").String()
			if config != "" {
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config)
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
