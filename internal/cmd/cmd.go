package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"framehub/internal/consts"
	"framehub/internal/controller/account"
	"framehub/internal/controller/users"
	"framehub/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						users.NewV1(),
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.AuthMiddleware(consts.User))
						group.Bind(
							account.NewV1(),
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
