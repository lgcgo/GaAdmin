// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AuthAccess is the golang structure of table auth_access for DAO operations like Where/Data.
type AuthAccess struct {
	g.Meta `orm:"table:auth_access, do:true"`
	Type   interface{} // 类型
	RuleId interface{} // 规则ID
}
