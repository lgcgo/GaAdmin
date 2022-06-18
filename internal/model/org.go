package model

import "GaAdmin/internal/model/entity"

// 组织管理
type OrgCreateInput struct {
	Name            string
	Phone           string
	Address         string
	CertificatesUrl string
	CertificatesNo  string
}
type OrgUpdateInput struct {
	OrgId           uint
	Name            string
	Phone           string
	Address         string
	CertificatesUrl string
	CertificatesNo  string
}
type OrgPageOutput struct {
	List []*entity.Org
	Pager
}

// 成员管理
type OrgMemberCreateInput struct {
	UserId       uint
	OrgId        uint
	Realname     string
	InitPassword string
	No           string
}
type OrgMemberUpdateInput struct {
	MemberId     uint
	Realname     string
	InitPassword string
	No           string
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
