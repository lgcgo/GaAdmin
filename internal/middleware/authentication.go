package middleware

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"GaAdmin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 权限认证中间件
func Authentication(r *ghttp.Request) {
	var (
		ser       = service.Oauth()
		claims    g.Map
		err       error
		issueRole string
		ok        bool
	)

	// 验证授权
	if claims, err = ser.ValidAuthorization(r); err != nil {
		response.JsonErrorExit(r, "11003", "invalid token")
	}
	// 从签名中获取用户角色
	issueRole = claims["isr"].(string)
	// 验证权限
	if ok, err = ser.CheckPath(r, issueRole); err != nil {
		response.JsonErrorExit(r, "-1", "system busy")
	}
	if !ok {
		response.JsonErrorExit(r, "11001", "no permission")
	}

	var (
		user    *entity.User
		isAdmin bool
	)

	// 取用户实体，先尝试会话获取
	if user = service.Session().GetUser(r.Context()); user == nil {
		user, err = service.User().GetUserByUuid(r.Context(), claims["sub"].(string))
		if err != nil {
			response.JsonErrorExit(r, "-1", "system busy")
		}
	}

	// 设置上下文
	isAdmin = true
	service.Context().SetUser(r.Context(), &model.ContextUser{
		Id:       user.Id,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		IsAdmin:  isAdmin,
	})

	r.Middleware.Next()
}
