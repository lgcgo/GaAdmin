package middleware

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"GaAdmin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// 权限认证中间件
func Authentication(r *ghttp.Request) {
	var (
		ser        = service.Oauth()
		claims     g.Map
		err        error
		issueRoles []string
		ok         bool
	)

	// 验证授权
	if claims, err = ser.ValidAuthorization(r); err != nil {
		response.JsonErrorExit(r, "11003", "invalid token")
	}
	// 从签名中获取用户角色
	issueRoles = gconv.Strings(claims["isr"])
	// 验证权限
	if ok, err = ser.CheckPath(r, issueRoles); err != nil {
		response.JsonErrorExit(r, "-1", "system busy")
	}
	if !ok {
		response.JsonErrorExit(r, "11001", "no permission")
	}

	var (
		user    *entity.User
		isAdmin bool
	)

	// 保持用户会话
	if user = service.Session().GetUser(r.Context()); user == nil {
		// 自动更新上线
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
