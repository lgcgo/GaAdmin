package model

import "GaAdmin/internal/model/entity"

/**
* 部门管理
**/
type OrgDepartmentCreateInput struct {
	ParentId uint
	Title    string
	Weigh    uint
}

type OrgDepartmentUpdateInput struct {
	DepartmentId uint
	ParentId     uint
	Title        string
	Weigh        uint
}

/**
* 成员管理
**/
type OrgMemberCreateInput struct {
	UserId       uint
	Realname     string
	InitPassword string
	No           string
}
type OrgMemberUpdateInput struct {
	MemberId     uint
	Realname     string
	InitPassword string
	No           string
	Status       string
}
type OrgMemberPageOutput struct {
	List []*entity.OrgMember
	Pager
}
type OrgMemberSignNoInput struct {
	Passport string
	Password string
	Captcha  string
}
