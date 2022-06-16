// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthMenu is the golang structure for table auth_menu.
type AuthMenu struct {
	Id       uint        `json:"id"       ` // ID
	ParentId uint        `json:"parentId" ` // 父ID
	Title    string      `json:"title"    ` // 标题
	Remark   string      `json:"remark"   ` // 备注
	Status   string      `json:"status"   ` // 状态
	Weigh    int         `json:"weigh"    ` // 权重
	CreateAt *gtime.Time `json:"createAt" ` // 创建日期
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新日期
}
