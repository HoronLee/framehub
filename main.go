package main

import (
	_ "framehub/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"framehub/internal/cmd"
)

func main() {
	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")
	cmd.Main.Run(gctx.GetInitCtx())
}
