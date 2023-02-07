package cronjob

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:        "gf2-demo-job",
		Brief:       "",
		Description: "",
		Usage:       "",
		Arguments: []gcmd.Argument{
			{},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Print("gf2-demo-job")

			return
		},
	}
)
