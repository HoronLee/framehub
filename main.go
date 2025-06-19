package main

import (
	"framehub/internal/logic/middleware"
	"framehub/internal/logic/users"
	_ "framehub/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	"framehub/internal/cmd"
)

// Gctx is the global context for the application.
var Gctx = gctx.GetInitCtx()

func main() {
	initLogger()
	initConfigs()
	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")
	cmd.Main.Run(Gctx)
}

func initLogger() {
	logConf := g.Cfg().MustGet(Gctx, "logger").MapStrAny()
	glog.SetPath(logConf["path"].(string))
	glog.SetFile(logConf["file"].(string))
	glog.SetPrefix(logConf["prefix"].(string))
	glog.SetLevelStr(logConf["level"].(string))
}

func initConfigs() {
	users.InitJwtConfig(Gctx)
	middleware.InitMwConfig(Gctx)
}
