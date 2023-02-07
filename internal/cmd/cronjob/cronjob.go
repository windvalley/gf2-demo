package cronjob

import (
	"context"
	"fmt"
	"gf2-demo/utility"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "gf2-demo-job",
		Brief:       "",
		Description: "",
		Usage:       "",
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

			fmt.Printf("gf2-demo-job\n")

			return
		},
	}
)
