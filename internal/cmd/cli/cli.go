package cli

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"gf2-demo/internal/consts"
	"gf2-demo/utility"
)

var (
	Main = gcmd.Command{
		Name:        "gf2-demo-cli",
		Brief:       "A command-line tool demo",
		Description: "A command-line tool demo using GoFrame V2",
		Usage:       "gf2-demo-cli [OPTION]",
		Examples: `
			Dev:
				./gf2-demo-cli

			Test:
				./gf2-demo-cli -c config.test.yaml
				or 
				GF_GCFG_FILE=config.test.yaml ./gf2-demo-cli

			Prod:
				./gf2-demo-cli -c config.prod.yaml
				or 
				GF_GCFG_FILE=config.prod.yaml ./gf2-demo-cli`,
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
				return nil
			}

			config := parser.GetOpt("config").String()
			if config != "" {
				//nolint: forcetypeassert
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config)
			}

			// json格式日志
			logFormat, err := g.Cfg().Get(ctx, "logger.cli.format")
			if err == nil {
				if logFormat.String() == "json" {
					glog.SetDefaultHandler(glog.HandlerJson)
				}
			}

			// 显示打印错误的文件行号
			g.Log(consts.CliLoggerName).SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)

			// 查看使用的配置文件是哪个
			configFile := g.Cfg().GetAdapter()
			g.Log(consts.CliLoggerName).Debugf(ctx, "use config file: %+v", configFile)

			// ****************** 以下部分为业务逻辑

			g.Log(consts.CliLoggerName).Info(ctx, "foo")

			g.Log(consts.CliLoggerName).Error(ctx, errors.New("bar"))

			return nil
		},
	}
)
