package cmd

import (
	"GaAdmin/internal/controller"
	"GaAdmin/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.CORS,     // 允许跨域
					middleware.Context,  // 初始化上下文对象
					middleware.Response, // 默认响应
				)
				group.Bind(
					controller.AuthRule,
					controller.AuthMenu,
					controller.Org,
					controller.OrgMember,
					controller.User,
					controller.UserGroup,
					controller.UserGroupAccess,
				)
			})
			s.Run()
			return nil
		},
	}
)
