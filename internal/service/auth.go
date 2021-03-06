// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type IAuth interface {
	InitRbac() error
	Authorization(subject string, role string) (*model.TokenOutput, error)
	RefreshAuthorization(ticket string) (*model.TokenOutput, error)
	VerifyToken(ticket string) (g.Map, error)
	VerifyRequest(path, method, role string) error
	CreateMenu(ctx context.Context, in *model.AuthMenuCreateInput) (*entity.AuthMenu, error)
	GetMenu(ctx context.Context, menuId uint) (*entity.AuthMenu, error)
	GetMenus(ctx context.Context, menuIds []uint) ([]*entity.AuthMenu, error)
	GetAllMenu(ctx context.Context) ([]*entity.AuthMenu, error)
	UpdateMenu(ctx context.Context, in *model.AuthMenuUpdateInput) (*entity.AuthMenu, error)
	DeleteMenu(ctx context.Context, menuId uint) error
	GetMenuTreeData(ctx context.Context) (*model.TreeDataOutput, error)
	GetMenuChildrenIds(ctx context.Context, menuId uint) ([]uint, error)
	CreateRole(ctx context.Context, in *model.AuthRoleCreateInput) (*entity.AuthRole, error)
	GetRole(ctx context.Context, roleId uint) (*entity.AuthRole, error)
	UpdateRole(ctx context.Context, in *model.AuthRoleUpdateInput) (*entity.AuthRole, error)
	DeleteRole(ctx context.Context, roleId uint) error
	GetAllRole(ctx context.Context) ([]*entity.AuthRole, error)
	GetRoleTreeData(ctx context.Context) (*model.TreeDataOutput, error)
	GetRoleName(ctx context.Context, roleID uint) (string, error)
	GetRoleChildrenIDs(ctx context.Context, roleId uint) ([]uint, error)
	CheckRoleIds(ctx context.Context, roleIds []uint) ([]uint, error)
	SetupRoleAccess(ctx context.Context, roleId uint, auth_rule_ids []uint) error
	DeleteRoleAccessByRuleID(ctx context.Context, ruleId uint) error
	GetAllRoleAccess(ctx context.Context) ([]*entity.AuthRoleAccess, error)
	CreateRule(ctx context.Context, in *model.AuthRuleCreateInput) (*entity.AuthRule, error)
	GetRule(ctx context.Context, nodeId uint) (*entity.AuthRule, error)
	GetRules(ctx context.Context, ruleIds []uint) ([]*entity.AuthRule, error)
	UpdateRule(ctx context.Context, in *model.AuthRuleUpdateInput) (*entity.AuthRule, error)
	DeleteRule(ctx context.Context, ruleId uint) error
	CheckRuleIds(ctx context.Context, ruleIds []uint) error
}

var localAuth IAuth

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
