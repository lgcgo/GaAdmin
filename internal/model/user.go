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
	GroupIds []uint
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
	GroupIds []uint
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

/**
* 分组管理
**/
type UserGroupCreateInput struct {
	ParentId uint
	Name     string
	Title    string
}

type UserGroupUpdateInput struct {
	GroupId  uint
	ParentId uint
	Name     string
	Title    string
}
