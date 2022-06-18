// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrgMember is the golang structure of table org_member for DAO operations like Where/Data.
type OrgMember struct {
	g.Meta       `orm:"table:org_member, do:true"`
	Id           interface{} // ID
	UserId       interface{} // 用户ID
	OrgId        interface{} // 公司ID
	Realname     interface{} // 真实名称
	No           interface{} // 工号
	InitPassword interface{} //
	Status       interface{} // 状态
	CreateAt     *gtime.Time // 创建日期
	UpdateAt     *gtime.Time // 更新日期
}
