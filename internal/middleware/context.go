package middleware

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 初始化上下文对象
func Context(r *ghttp.Request) {
	var (
		customCtx *model.Context
	)
	// 初始化，务必最开始执行
	customCtx = &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	service.Context().Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
