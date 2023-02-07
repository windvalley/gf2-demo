package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "gf2-demo/internal/logic"
	_ "gf2-demo/internal/packed"

	"gf2-demo/internal/cmd/apiserver"
)

func main() {
	apiserver.Main.Run(gctx.New())
}
