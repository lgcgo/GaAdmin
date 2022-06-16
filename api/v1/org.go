package v1

import "github.com/gogf/gf/v2/frame/g"

/**
* 组织机构
**/
// 组织返回的数据项
type OrgResData struct {
	Id              uint   `json:"orgId"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	CertificatesUrl string `json:"certificatesUrl"`
	CertificatesNo  string `json:"certificatesNo"`
}

// 创建组织
type OrgCreateReq struct {
	g.Meta          `path:"/org" method:"post" tags:"OrgService" summary:"Create org"`
	Name            string `json:"name" v:"required|length:4,16"`
	Phone           string `json:"phone" v:"required"`
	Address         string `json:"address"`
	CertificatesUrl string `json:"certificatesUrl"`
	CertificatesNo  string `json:"certificatesNo"`
}
type OrgCreateRes struct {
	OrgResData
}

// 获取组织
type OrgGetReq struct {
	g.Meta `path:"/org" method:"get" tags:"OrgService" summary:"Get org"`
	OrgId  uint `json:"orgId" v:"required"`
}
type OrgGetRes struct {
	OrgResData
}

// 修改组织
type OrgUpdateReq struct {
	g.Meta          `path:"/org" method:"put" tags:"OrgService" summary:"Update org"`
	OrgId           uint   `json:"orgId" v:"required"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	CertificatesUrl string `json:"certificatesUrl"`
	CertificatesNo  string `json:"certificatesNo"`
}
type OrgUpdateRes struct {
	Id uint `json:"orgId"`
	OrgResData
}

// 删除组织
type OrgDeleteReq struct {
	g.Meta `path:"/org" method:"delete" tags:"OrgService" summary:"Delete org"`
	OrgId  uint `json:"orgId" v:"required"`
}
type OrgDeleteRes struct{}

// 获取组织列表
type OrgListReq struct {
	g.Meta `path:"/orgs" method:"get" tags:"OrgService" summary:"Get org list"`
	Page
}
type OrgListRes struct {
	List []*OrgResData `json:"list"`
	Pager
}

/**
* 组织成员
**/
// 返回的数据项
type OrgMemberResData struct {
	Id       uint   `json:"memberId"`              // 成员ID
	OrgId    uint   `json:"orgId" v:"required"`    // 公司ID
	Realname string `json:"realname" v:"required"` // 真实名称
	No       string `json:"no"`                    // 工号
	Status   string `json:"status"`                // 状态
}

// 创建成员
type OrgMemberCreateReq struct {
	g.Meta   `path:"/org/member" method:"post" tags:"OrgService" summary:"Create member"`
	OrgId    uint   `json:"orgId" v:"required"`    // 公司ID
	Realname string `json:"realname" v:"required"` // 真实名称
	No       string `json:"no"`                    // 工号
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
