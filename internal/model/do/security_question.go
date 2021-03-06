// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SecurityQuestion is the golang structure of table security_question for DAO operations like Where/Data.
type SecurityQuestion struct {
	g.Meta   `orm:"table:security_question, do:true"`
	Id       interface{} // ID
	Title    interface{} // 标题
	CreateAt *gtime.Time // 创建日期
}
