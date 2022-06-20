package v1

import "github.com/gogf/gf/v2/frame/g"

/**
* 基础权限
**/
// 设置基础权限
type AuthAccessSetupBasicReq struct {
	g.Meta  `path:"/auth/access/setup-basic" tags:"AuthService" method:"post" summary:"Setup basic access"`
	RuleIds []uint `json:"ruleIds" v:"required"`
}
type AuthAccessSetupBasicRes struct {
	RuleIds []uint `json:"ruleIds"`
}

// 设置限制权限（被禁用时仍拥有的权限）
type AuthAccessSetupLimitedReq struct {
	g.Meta  `path:"/auth/access/setup-limited" tags:"AuthService" method:"post" summary:"Setup limited access"`
	RuleIds []uint `json:"ruleIds" v:"required"`
}
type AuthAccessSetupLimitedRes struct {
	RuleIds []uint `json:"ruleIds"`
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
	ParentId uint   `json:"parentId" v:"required"`
	Title    string `json:"title" v:"required"`
	Remark   string `json:"remark"`
	Weigh    uint   `json:"weigh"`
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
	MenuId   uint   `json:"menuId" v:"required"`
	ParentId uint   `json:"parentId" v:"required"`
	Title    string `json:"title" v:"required"`
	Remark   string `json:"remark"`
	Weigh    uint   `json:"weigh"`
}
type AuthMenuUpdateRes struct {
	AuthMenuResData
}

// 删除权限菜单
type AuthMenuDeleteReq struct {
	g.Meta `path:"/auth/menu" tags:"AuthService" method:"delete" summary:"Delete menu"`
	MenuId uint `json:"menuId" v:"required"`
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
	MenuId    uint   `json:"menuId" v:"required"`
	Title     string `json:"title" v:"required"`
	Path      string `json:"path" v:"required"`
	Method    string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Condition string `json:"condition"`
	Remark    string `json:"remark"`
	Weigh     uint   `json:"weigh"`
}
type AuthRuleCreateRes struct {
	AuthRuleResData
}

// 获取权限节点
type AuthRuleGetReq struct {
	g.Meta `path:"/auth/rule" tags:"AuthService" method:"get" summary:"Get rule"`
	RuleId uint `json:"ruleId" v:"required"`
}
type AuthRuleGetRes struct {
	AuthRuleResData
}

// 更新权限节点
type AuthRuleUpdateReq struct {
	g.Meta    `path:"/auth/rule" tags:"AuthService" method:"put" summary:"Update rule"`
	RuleId    uint   `json:"ruleId" v:"required"`
	MenuId    uint   `json:"menuId" v:"required"`
	Title     string `json:"title" v:"required"`
	Path      string `json:"path" v:"required"`
	Method    string `json:"method" v:"required|in:GET,POST,PUT,DELETE,PATCH"`
	Condition string `json:"condition"`
	Remark    string `json:"remark"`
}
type AuthRuleUpdateRes struct {
	AuthRuleResData
}

// 删除权限节点
type AuthRuleDeleteReq struct {
	g.Meta `path:"/auth/rule" tags:"AuthService" method:"delete" summary:"Delete rule"`
	RuleId uint `json:"ruleId" v:"required"`
}
type AuthRuleDeleteRes struct {
}
