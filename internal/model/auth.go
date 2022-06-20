package model

/**
* 分组管理
**/
type AuthRoleCreateInput struct {
	ParentId uint
	Name     string
	Title    string
}

type AuthRoleUpdateInput struct {
	RoleId   uint
	ParentId uint
	Name     string
	Title    string
}

/**
* 菜单管理
**/
type AuthMenuCreateInput struct {
	ParentId uint
	Title    string
	Remark   string
	Weigh    uint
}

type AuthMenuUpdateInput struct {
	MenuId   uint
	ParentId uint
	Title    string
	Remark   string
	Weigh    uint
}

/**
* 规则管理
**/
type AuthRuleCreateInput struct {
	MenuId    uint
	Title     string
	Path      string
	Method    string
	Condition string
	Remark    string
	Weigh     uint
}

type AuthRuleUpdateInput struct {
	RuleId    uint
	MenuId    uint
	Title     string
	Path      string
	Method    string
	Condition string
	Remark    string
	Weigh     uint
}
