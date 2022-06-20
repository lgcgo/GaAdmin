package model

import "GaAdmin/internal/model/entity"

// 创建用户
type UserCreateInput struct {
	Account  string
	Password string
	Nickname string
	Avatar   string
	Mobile   string
	Email    string
}

// 修改用户
type UserUpdateInput struct {
	UserId   uint
	Account  string
	Password string
	Nickname string
	Avatar   string
	Mobile   string
	Email    string
}

// 用户分页
type UserPageOutput struct {
	List []*entity.User
	Pager
}

type UserSignPassportInput struct {
	Passport string
	Password string
	Captcha  string
}

type UserSignMobile struct {
	Mobile  string
	Captcha string
}
