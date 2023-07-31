package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/windvalley/gf2-demo/internal/logic"
	_ "github.com/windvalley/gf2-demo/internal/packed"

	"github.com/windvalley/gf2-demo/internal/cmd/cli"
)

func main() {
	cli.Main.Run(gctx.New())
}
