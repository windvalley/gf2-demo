package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	_ "gf2-demo/internal/logic"
	_ "gf2-demo/internal/packed"

	"gf2-demo/internal/cmd/apiserver"
)

func main() {
	apiserver.Main.Run(gctx.New())
}
