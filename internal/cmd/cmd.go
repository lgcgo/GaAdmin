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
					controller.UserSign,  // 用户登录
					controller.UserReset, // 重置密码
				)
				group.Middleware(middleware.Authentication)
				group.Bind(
					controller.AuthMenu,       // 权限菜单
					controller.AuthRole,       // 权限角色
					controller.AuthRoleAccess, // 角色授权
					controller.AuthRule,       // 权限规则
					controller.OrgDepartment,  // 组织部门
					controller.OrgMember,      // 组织成员
					controller.User,           // 用户管理
					controller.UserAccess,     // 用户授权
				)
			})
			s.Run()
			return nil
		},
	}
)
