package main

import (
	_ "gf2-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf2-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
