package cmd

import (
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
					middleware.Response, // 默认响应
				)
				group.Bind()
				// 权限认证路由
				// group.Middleware(middleware.Authentication)
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
