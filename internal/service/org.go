// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"context"
)

type IOrg interface {
	GetMemberByUuid(ctx context.Context, uuid string) (*entity.OrgMember, error)
	CreateMember(ctx context.Context, in *model.OrgMemberCreateInput) (*entity.OrgMember, error)
	GetMember(ctx context.Context, memberId uint) (*entity.OrgMember, error)
	UpdateMember(ctx context.Context, in *model.OrgMemberUpdateInput) (*entity.OrgMember, error)
	DeleteMember(ctx context.Context, id uint) error
	GetMemberPage(ctx context.Context, in *model.Page) (*model.OrgMemberPageOutput, error)
	IsMemberUuidAvailable(ctx context.Context, orgId uint, uuid string, notIds ...uint) (bool, error)
	IsMemberNoAvailable(ctx context.Context, orgId uint, no string, notIds ...uint) (bool, error)
	CreateOrg(ctx context.Context, in *model.OrgCreateInput) (*entity.Org, error)
	GetOrg(ctx context.Context, orgId uint) (*entity.Org, error)
	UpdateOrg(ctx context.Context, in *model.OrgUpdateInput) (*entity.Org, error)
	DeleteOrg(ctx context.Context, id uint) error
	GetOrgPage(ctx context.Context, in *model.Page) (*model.OrgPageOutput, error)
	IsNameAvailable(ctx context.Context, name string, notIds ...uint) (bool, error)
	IsCertificatesNoAvailable(ctx context.Context, certificatesNo string, notIds ...uint) (bool, error)
}

var localOrg IOrg

func Org() IOrg {
	if localOrg == nil {
		panic("implement not found for interface IOrg, forgot register?")
	}
	return localOrg
}

func RegisterOrg(i IOrg) {
	localOrg = i
}
