// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRoleAccess is the golang structure of table auth_role_access for DAO operations like Where/Data.
type AuthRoleAccess struct {
	g.Meta `orm:"table:auth_role_access, do:true"`
	RoleId interface{} // 角色ID
	RuleId interface{} // 规则ID
}
