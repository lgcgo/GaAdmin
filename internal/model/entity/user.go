// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id           uint        `json:"id"           ` // ID
	Uuid         string      `json:"uuid"         ` // 唯一ID
	Account      string      `json:"account"      ` // 账号
	Password     string      `json:"password"     ` // 密码
	Salt         string      `json:"salt"         ` // 密码盐
	Nickname     string      `json:"nickname"     ` // 昵称
	Avatar       string      `json:"avatar"       ` // 头像
	Mobile       string      `json:"mobile"       ` // 手机号
	Email        string      `json:"email"        ` // 电子邮箱
	GroupIds     string      `json:"groupIds"     ` // 用户组ID集
	Loginfailure uint        `json:"loginfailure" ` // 失败次数
	Loginip      string      `json:"loginip"      ` // 登录IP
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  ` // 登录日期
	Status       string      `json:"status"       ` // 状态
	CreateAt     *gtime.Time `json:"createAt"     ` // 创建日期
	UpdateAt     *gtime.Time `json:"updateAt"     ` // 更新日期
}
