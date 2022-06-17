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
					controller.UserSign, // 用户登录
				)
				group.Middleware(middleware.Authentication)
				group.Bind(
					controller.AuthRule,        // 权限规则
					controller.AuthMenu,        // 权限菜单
					controller.Org,             // 组织机构
					controller.OrgMember,       // 组织成员
					controller.User,            // 用户
					controller.UserGroup,       // 用户分组
					controller.UserGroupAccess, // 用户分组权限
				)
			})
			s.Run()
			return nil
		},
	}
)
