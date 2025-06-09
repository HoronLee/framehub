package main

import (
	_ "framehub/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"framehub/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
