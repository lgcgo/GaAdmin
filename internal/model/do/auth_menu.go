// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthMenu is the golang structure of table auth_menu for DAO operations like Where/Data.
type AuthMenu struct {
	g.Meta   `orm:"table:auth_menu, do:true"`
	Id       interface{} // ID
	ParentId interface{} // 父ID
	Title    interface{} // 标题
	Remark   interface{} // 备注
	Status   interface{} // 状态
	Weigh    interface{} // 权重
	CreateAt *gtime.Time // 创建日期
	UpdateAt *gtime.Time // 更新日期
}
