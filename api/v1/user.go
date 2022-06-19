package v1

import "github.com/gogf/gf/v2/frame/g"

// 用户返回的数据项
type UserResData struct {
	Id           uint   `json:"userId"`       //用户Id
	Uuid         string `json:"uuid"`         // 唯一ID
	Account      string `json:"account"`      // 账号
	Nickname     string `json:"nickname"`     // 昵称
	Avatar       string `json:"avatar"`       // 头像
	Mobile       string `json:"mobile"`       // 手机号
	Email        string `json:"email"`        // 电子邮箱
	Loginfailure uint   `json:"loginfailure"` // 失败次数
	Loginip      string `json:"loginip"`      // 登录IP
	LastLoginAt  string `json:"lastLoginAt"`  // 登录日期
	Status       string `json:"status"`       // 状态
	CreateAt     string `json:"createAt"`     // 创建日期
	UpdateAt     string `json:"updateAt"`     // 更新日期
}

// 创建用户（用于后台，账户密码必填）
type UserCreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"UserService" summary:"Create user"`
	Account  string `json:"account" v:"required"`  // 账号
	Password string `json:"password" v:"required"` // 密码
	Nickname string `json:"nickname" v:"required"` // 昵称
	Avatar   string `json:"avatar"`                // 头像
	Mobile   string `json:"mobile"`                // 手机号
	Email    string `json:"email"`                 // 电子邮箱
}
type UserCreateRes struct {
	UserResData
}

// 获取用户
type UserGetReq struct {
	g.Meta `path:"/user" method:"get" tags:"UserService" summary:"Get user"`
	UserId uint `json:"userId" v:"required"`
}
type UserGetRes struct {
	UserResData
}

// 修改用户
type UserUpdateReq struct {
	g.Meta   `path:"/user" method:"put" tags:"UserService" summary:"Update user"`
	UserId   uint   `json:"userId" v:"required"`
	Account  string `json:"account"`  // 账号
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	Mobile   string `json:"mobile"`   // 手机号
	Email    string `json:"email"`    // 电子邮箱
}
type UserUpdateRes struct {
	UserResData
}

// 删除用户
type UserDeleteReq struct {
	g.Meta `path:"/user" method:"delete" tags:"UserService" summary:"Delete user"`
	UserId uint `json:"userId" v:"required"`
}
type UserDeleteRes struct{}

// 用户列表
type UserListReq struct {
	g.Meta `path:"/users" method:"get" tags:"UserService" summary:"Get user list"`
	Page
}
type UserListRes struct {
	List []*UserResData `json:"list"`
	Pager
}

/**
*  用户分组
**/
// 分组返回数据
type UserGroupResData struct {
	Id       uint   `json:"groupId"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Title    string `json:"title"`
}

// 创建分组
type UserGroupCreateReq struct {
	g.Meta   `path:"/user/group" tags:"UserService" method:"post" summary:"Create group"`
	ParentId uint   `json:"parentId" v:"required"`
	Name     string `json:"name" v:"required"`
	Title    string `json:"title" v:"required"`
}
type UserGroupCreateRes struct {
	UserGroupResData
}

// 获取分组
type UserGroupGetReq struct {
	g.Meta  `path:"/user/group" tags:"UserService" method:"get" summary:"Get group"`
	GroupId uint `json:"groupId" v:"required"`
}
type UserGroupGetRes struct {
	UserGroupResData
}

// 更新分组
type UserGroupUpdateReq struct {
	g.Meta   `path:"/user/group" tags:"UserService" method:"put" summary:"Update group"`
	GroupId  uint   `json:"groupId" v:"required"`
	ParentId uint   `json:"parentId" v:"required"`
	Name     string `json:"name" v:"required"`
	Title    string `json:"title" v:"required"`
}
type UserGroupUpdateRes struct {
	UserGroupResData
}

// 删除分组
type UserGroupDeleteReq struct {
	g.Meta  `path:"/user/group" tags:"UserService" method:"delete" summary:"Delete group"`
	GroupId uint `json:"groupId" v:"required"`
}
type UserGroupDeleteRes struct {
}

// 获取分组树
type UserGroupTreeReq struct {
	g.Meta `path:"/user/group-tree" tags:"UserService" method:"get" summary:"Get group tree"`
}
type UserGroupTreeRes struct {
	TreeResData
}

// 设置分组权限
type UserGroupAccessSetupReq struct {
	g.Meta      `path:"/user/group-access" tags:"UserService" method:"post" summary:"Setup group access"`
	GroupId     uint   `json:"groupId" v:"required"`
	AuthRuleIds []uint `json:"authRuleIds" v:"required"`
}
type UserGroupAccessSetupRes struct {
	GroupId     uint   `json:"groupId"`
	AuthRuleIds []uint `json:"authRuleIds"`
}

/**
* 用户角色
**/
type UserRolesSetupReq struct {
	g.Meta   `path:"/user/roles" tags:"UserService" method:"post" summary:"Setup roles"`
	UserId   uint   `json:"userId" v:"required"`
	GroupIds []uint `json:"groupIds"`
}
type UserRolesSetupRes struct {
	UserId   uint   `json:"userId"`
	GroupIds []uint `json:"groupIds"`
}

/**
*  用户认证
**/
// 授权返回
type TokenResData struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

// 刷新token
type UserSignRefreshReq struct {
	g.Meta       `path:"/user/refresh-token" tags:"UserService" method:"post" summary:"Sign refresh"`
	RefreshToken string `json:"refreshToken" v:"required"`
}
type UserSignRefreshRes struct {
	TokenResData
}

// 注册用户
// - 至少保证一种登录方式原则
// - 账号|手机号|邮箱其中一个必填
// - 当存在账号，则密码必填
// - 当手存在机号|邮箱，则验证码必填
type UserSignUpReq struct {
	g.Meta   `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up"`
	Account  string `json:"account" v:"required-without-all:Mobile,Email|length:4,32"` // 账号
	Password string `json:"password" v:"required-with:Account|length:6,32"`            // 密码
	Nickname string `json:"nickname" v:"required|length:4,32"`                         // 昵称
	Mobile   string `json:"mobile" v:"required-without-all:Account,Email|phone"`       // 手机号
	Captcha  string `json:"captcha" v:"required-with:Mobile,Email|length:4,8"`         // 验证码
	Email    string `json:"email" v:"required-without-all:Account,Mobile|email"`       // 电子邮箱
	Avatar   string `json:"avatar"`                                                    // 头像
}
type UserSignUpRes struct {
	TokenResData
}

// 账户|手机号|邮箱 + 密码 登录
// 登录次数超过后，服务要求 Captcha 验证
type UserSignPassportReq struct {
	g.Meta   `path:"/user/sign-passport" tags:"UserService" method:"post" summary:"Sign in passport"`
	Passport string `json:"passport" v:"required|length:4,32"` // 账户|手机号|邮箱
	Password string `json:"password" v:"required|length:6,32"` // 密码
	Captcha  string `json:"captcha" v:"length:4,8"`            // 验证码
}
type UserSignPassportRes struct {
	TokenResData
}

// 手机号 + 短信验证码 登录
type UserSignMobileReq struct {
	g.Meta  `path:"/user/sign-mobile" tags:"UserService" method:"post" summary:"Sign in mobile"`
	Mobile  string `json:"mobile" v:"required|phone"`       // 手机号
	Captcha string `json:"captcha" v:"required:length:4,8"` // 验证码
}
type UserSignMobileRes struct {
	TokenResData
}

// 注销登录
type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"put" tags:"UserService" summary:"Sign out"`
}
type UserSignOutRes struct{}

// 手机 密码重置
type UserResetMobileReq struct {
	g.Meta  `path:"/user/password/reset-mobile" tags:"UserService" method:"post" summary:"Reset password mobile"`
	Email   string `json:"email" v:"required|email"`
	Captcha string `json:"captcha" v:"required|length:4,8"`
}
type UserResetMobileRes struct {
}

// 邮件 密码重置
type UserResetEmailReq struct {
	g.Meta  `path:"/user/password/reset-email" tags:"UserService" method:"post" summary:"Reset password email"`
	Email   string `json:"email" v:"required|email"`
	Captcha string `json:"captcha" v:"required|length:4,8"`
}
type UserResetEmailRes struct {
}

// 问答 密码重置
type UserResetAnswer struct {
	QuestionId uint   // 问题ID
	AnswerText string // 回答内容
}
type UserResetQuestionReq struct {
	g.Meta  `path:"/user/password/reset-question" tags:"UserService" method:"post" summary:"Reset password question"`
	Answers []UserResetAnswer `json:"answers" v:"required"`
}
type UserResetQuestionRes struct {
}
