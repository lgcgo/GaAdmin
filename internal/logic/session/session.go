package session

import (
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"
)

type sSession struct{}

const (
	sessionKeyUser = "SessionKeyUser" // 用户信息存放在Session中的Key
)

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// 设置用户Session.
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.Context().Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := service.Context().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.User{}
}

// 删除用户Session。
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
