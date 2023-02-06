package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "gf2-api/internal/logic"
	_ "gf2-api/internal/packed"

	"gf2-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
