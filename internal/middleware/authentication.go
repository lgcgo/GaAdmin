package middleware

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"GaAdmin/utility/response"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 权限认证中间件
func Authentication(r *ghttp.Request) {
	var (
		auth   = service.Auth()
		claims g.Map
		err    error
		ok     bool
	)

	if err = auth.InitRbac(); err != nil {
		response.JsonErrorExit(r, "-1", "system busy")
	}

	// Header传值 Authorization: Bearer <token>
	if r.Header.Get("Authorization") == "" {
		panic("headers authorization not exists")
	}
	strArr := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	// 支持Bearer方案
	if strArr[0] != "Bearer" {
		panic("authorization scheme not support")
	}

	// 获取Token票据
	tokenTicket := strArr[1]

	// 验证授权
	if claims, err = auth.VerifyToken(tokenTicket); err != nil {
		response.JsonErrorExit(r, "11003", "invalid token")
	}

	// 从签名中获取用户角色
	role := claims["isr"].(string)
	// 验证权限
	if err = auth.VerifyRequest(r.URL.Path, r.Method, role); err != nil {
		response.JsonErrorExit(r, "-1", "system busy")
	}
	if !ok {
		response.JsonErrorExit(r, "11001", "no permission")
	}

	var (
		user *entity.User
	)

	// 获取用户实体，优先会话获取
	user = service.Session().GetUser(r.Context())
	if user.Id == 0 {
		user, err = service.User().GetUserByUuid(r.Context(), claims["sub"].(string))
		if err != nil {
			response.JsonErrorExit(r, "-1", "system busy")
		}
	}

	// 设置上下文用户
	service.Context().SetUser(r.Context(), &model.ContextUser{
		Id:       user.Id,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	})

	r.Middleware.Next()
}
