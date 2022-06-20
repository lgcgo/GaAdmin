package v1

import "github.com/gogf/gf/v2/frame/g"

/**
*  组织部门
**/
// 部门返回数据
type OrgDepartmentResData struct {
	Id       uint   `json:"departmentId"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Weigh    uint   `json:"weigh"`
}

// 创建部门
type OrgDepartmentCreateReq struct {
	g.Meta   `path:"/org/department" tags:"OrgService" method:"post" summary:"Create department"`
	ParentId uint   `json:"parentId" v:"required"`
	Name     string `json:"name" v:"required"`
	Title    string `json:"title" v:"required"`
	Weigh    uint   `json:"weigh"`
}
type OrgDepartmentCreateRes struct {
	OrgDepartmentResData
}

// 获取部门
type OrgDepartmentGetReq struct {
	g.Meta       `path:"/org/department" tags:"OrgService" method:"get" summary:"Get department"`
	DepartmentId uint `json:"departmentId" v:"required"`
}
type OrgDepartmentGetRes struct {
	OrgDepartmentResData
}

// 更新部门
type OrgDepartmentUpdateReq struct {
	g.Meta       `path:"/org/department" tags:"OrgService" method:"put" summary:"Update department"`
	DepartmentId uint   `json:"departmentId" v:"required"`
	ParentId     uint   `json:"parentId" v:"required"`
	Title        string `json:"title" v:"required"`
	Weigh        uint   `json:"weigh"`
}
type OrgDepartmentUpdateRes struct {
	OrgDepartmentResData
}

// 删除部门
type OrgDepartmentDeleteReq struct {
	g.Meta       `path:"/org/department" tags:"OrgService" method:"delete" summary:"Delete department"`
	DepartmentId uint `json:"departmentId" v:"required"`
}
type OrgDepartmentDeleteRes struct {
}

// 获取部门树
type OrgDepartmentTreeReq struct {
	g.Meta `path:"/org/department-tree" tags:"OrgService" method:"get" summary:"Get department tree"`
}
type OrgDepartmentTreeRes struct {
	TreeResData
}

/**
* 组织成员
**/
// 返回的数据项
type OrgMemberResData struct {
	Id           uint   `json:"memberId"`              // 成员ID
	Realname     string `json:"realname" v:"required"` // 真实名称
	InitPassword string `json:"initPassword"`          // 初始密码
	No           string `json:"no"`                    // 工号
	Status       string `json:"status"`                // 状态
}

// 创建成员
type OrgMemberCreateReq struct {
	g.Meta       `path:"/org/member" method:"post" tags:"OrgService" summary:"Create member"`
	Realname     string `json:"realname" v:"required"` // 真实名称
	InitPassword string `json:"initPassword"`          // 初始密码
	No           string `json:"no"`                    // 工号
}
type OrgMemberCreateRes struct {
	OrgMemberResData
}

// 获取成员
type OrgMemberGetReq struct {
	g.Meta   `path:"/org/member" method:"get" tags:"OrgService" summary:"Get member"`
	MemberId uint `json:"memberId" v:"required"` // 成员ID
}
type OrgMemberGetRes struct {
	OrgMemberResData
}

// 修改成员
type OrgMemberUpdateReq struct {
	g.Meta   `path:"/org/member" method:"put" tags:"OrgService" summary:"Update member"`
	MemberId uint   `json:"memberId" v:"required"` // 成员ID
	Realname string `json:"realname" v:"required"` // 真实名称
	No       string `json:"no"`                    // 工号
}
type OrgMemberUpdateRes struct {
	OrgMemberResData
}

// 删除成员
type OrgMemberDeleteReq struct {
	g.Meta   `path:"/org/member" method:"delete" tags:"OrgService" summary:"Delete member"`
	MemberId uint `json:"memberId" v:"required"` // 成员Id
}
type OrgMemberDeleteRes struct {
}

// 获取成员列表
type OrgMemberListReq struct {
	g.Meta `path:"/org/members" method:"get" tags:"OrgService" summary:"Get member list"`
	Page
}
type OrgMemberListRes struct {
	List []*OrgMemberResData `json:"list"`
	Pager
}
