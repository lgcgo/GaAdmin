package v1

import "github.com/gogf/gf/v2/frame/g"

/**
*  权限角色
**/
// 返回角色数据
type AuthRoleResData struct {
	Id       uint   `json:"roleId"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Title    string `json:"title"`
}

// 创建角色
type AuthRoleCreateReq struct {
	g.Meta   `path:"/auth/role" tags:"AuthService" method:"post" summary:"Create group"`
	ParentId uint   `json:"parentId" v:"required|integer"`
	Name     string `json:"name" v:"required|length:4,16"`
	Title    string `json:"title" v:"required|length:4,16"`
}
type AuthRoleCreateRes struct {
	AuthRoleResData
}

// 获取角色
type AuthRoleGetReq struct {
	g.Meta `path:"/auth/role" tags:"AuthService" method:"get" summary:"Get group"`
	RoleId uint `json:"roleId" v:"required|integer"`
}
type AuthRoleGetRes struct {
	AuthRoleResData
}

// 更新角色
type AuthRoleUpdateReq struct {
	g.Meta   `path:"/auth/role" tags:"AuthService" method:"put" summary:"Update group"`
	RoleId   uint   `json:"roleId" v:"required|integer"`
	ParentId uint   `json:"parentId" v:"required|integer"`
	Name     string `json:"name" v:"required|length:4,16"`
	Title    string `json:"title" v:"required|length:4,16"`
}
type AuthRoleUpdateRes struct {
	AuthRoleResData
}

// 删除角色
type AuthRoleDeleteReq struct {
	g.Meta `path:"/auth/role" tags:"AuthService" method:"delete" summary:"Delete group"`
	RoleId uint `json:"roleId" v:"required|integer"`
}
type AuthRoleDeleteRes struct {
}

// 获取角色树
type AuthRoleTreeReq struct {
	g.Meta `path:"/auth/role-tree" tags:"AuthService" method:"get" summary:"Get group tree"`
}
type AuthRoleTreeRes struct {
	TreeResData
}

// 设置角色权限
type AuthRoleAccessSetupReq struct {
	g.Meta      `path:"/auth/role-access" tags:"AuthService" method:"post" summary:"Setup group access"`
	RoleId      uint   `json:"roleId" v:"required|integer"`
	AuthRuleIds []uint `json:"authRuleIds"`
}
type AuthRoleAccessSetupRes struct {
	RoleId      uint   `json:"roleId"`
	AuthRuleIds []uint `json:"authRuleIds"`
}

/**
*  权限菜单
**/
// 权限菜单返回数据
type AuthMenuResData struct {
	Id       uint   `json:"menuId"`
	ParentId uint   `json:"parentId"`
	Title    string `json:"title"`
	Remark   string `json:"remark"`
	Weigh    uint   `json:"weigh"`
}

// 创建权限菜单
type AuthMenuCreateReq struct {
	g.Meta   `path:"/auth/menu" tags:"AuthService" method:"post" summary:"Create menu"`
	ParentId uint   `json:"parentId" v:"required|integer"`
	Title    string `json:"title" v:"required|length:4,16"`
	Remark   string `json:"remark" v:"max-length:32"`
	Weigh    uint   `json:"weigh" v:"between:0,9999"`
}
type AuthMenuCreateRes struct {
	AuthMenuResData
}

// 获取权限菜单
type AuthMenuGetReq struct {
	g.Meta `path:"/auth/menu" tags:"AuthService" method:"get" summary:"Get menu"`
	MenuId uint `json:"menuId" v:"required"`
}
type AuthMenuGetRes struct {
	AuthMenuResData
}

// 更新权限菜单
type AuthMenuUpdateReq struct {
	g.Meta   `path:"/auth/menu" tags:"AuthService" method:"put" summary:"Update menu"`
	MenuId   uint   `json:"menuId" v:"required|integer"`
	ParentId uint   `json:"parentId" v:"required|integer"`
	Title    string `json:"title" v:"required|length:4,16"`
	Remark   string `json:"remark" v:"max-length:32"`
	Weigh    uint   `json:"weigh" v:"between:0,9999"`
}
type AuthMenuUpdateRes struct {
	AuthMenuResData
}

// 删除权限菜单
type AuthMenuDeleteReq struct {
	g.Meta `path:"/auth/menu" tags:"AuthService" method:"delete" summary:"Delete menu"`
	MenuId uint `json:"menuId" v:"required|integer"`
}
type AuthMenuDeleteRes struct {
}

type AuthMenuTreeReq struct {
	g.Meta `path:"/auth/menu-tree" tags:"AuthService" method:"get" summary:"Get menu tree"`
}
type AuthMenuTreeRes struct {
	TreeResData
}

/**
*  权限节点
**/
// 权限节点返回数据
type AuthRuleResData struct {
	Id        uint   `json:"ruleId"`
	MenuId    uint   `json:"menuId"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Condition string `json:"condition"`
	Remark    string `json:"remark"`
	Weigh     uint   `json:"weigh"`
}

// 创建权限节点
type AuthRuleCreateReq struct {
	g.Meta    `path:"/auth/rule" tags:"AuthService" method:"post" summary:"Create rule"`
	MenuId    uint   `json:"menuId" v:"required|integer"`
	Title     string `json:"title" v:"required|length:4,16"`
	Path      string `json:"path" v:"required|length:2,64"`
	Method    string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Condition string `json:"condition" v:"max-length:512"`
	Remark    string `json:"remark" v:"max-length:32"`
	Weigh     uint   `json:"weigh" v:"integer|between:0,9999"`
}
type AuthRuleCreateRes struct {
	AuthRuleResData
}

// 获取权限节点
type AuthRuleGetReq struct {
	g.Meta `path:"/auth/rule" tags:"AuthService" method:"get" summary:"Get rule"`
	RuleId uint `json:"ruleId" v:"required|integer"`
}
type AuthRuleGetRes struct {
	AuthRuleResData
}

// 更新权限节点
type AuthRuleUpdateReq struct {
	g.Meta    `path:"/auth/rule" tags:"AuthService" method:"put" summary:"Update rule"`
	RuleId    uint   `json:"ruleId" v:"required|integer"`
	MenuId    uint   `json:"menuId" v:"required|integer"`
	Title     string `json:"title" v:"required|length:4,16"`
	Path      string `json:"path" v:"required|length:2,64"`
	Method    string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Condition string `json:"condition" v:"max-length:512"`
	Remark    string `json:"remark" v:"max-length:32"`
	Weigh     uint   `json:"weigh" v:"integer|between:0,9999"`
}
type AuthRuleUpdateRes struct {
	AuthRuleResData
}

// 删除权限节点
type AuthRuleDeleteReq struct {
	g.Meta `path:"/auth/rule" tags:"AuthService" method:"delete" summary:"Delete rule"`
	RuleId uint `json:"ruleId" v:"required|integer"`
}
type AuthRuleDeleteRes struct {
}
