// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrgMember is the golang structure for table org_member.
type OrgMember struct {
	Id           uint        `json:"id"           ` // ID
	Uuid         string      `json:"uuid"         ` // 唯一ID
	OrgId        uint        `json:"orgId"        ` // 公司ID
	Realname     string      `json:"realname"     ` // 真实名称
	No           string      `json:"no"           ` // 工号
	InitPassword string      `json:"initPassword" ` //
	Status       string      `json:"status"       ` // 状态
	CreateAt     *gtime.Time `json:"createAt"     ` // 创建日期
	UpdateAt     *gtime.Time `json:"updateAt"     ` // 更新日期
}
