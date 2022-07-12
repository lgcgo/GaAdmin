package auth

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 设置角色权限
func (s *sAuth) SetupRoleAccess(ctx context.Context, roleId uint, auth_rule_ids []uint) error {
	var (
		err  error
		role *entity.AuthRole
		data []*do.AuthRoleAccess
	)

	// 检测角色
	if role, err = s.GetRole(ctx, roleId); err != nil {
		return err
	}
	if role == nil {
		return gerror.Newf("role is not exists: %d", roleId)
	}
	// 检测权限ID集
	if err = service.Auth().CheckRuleIds(ctx, auth_rule_ids); err != nil {
		return err
	}
	// 组装新增数据
	for _, v := range auth_rule_ids {
		data = append(data, &do.AuthRoleAccess{
			RoleId: roleId,
			RuleId: v,
		})
	}
	if err = dao.AuthRoleAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 移除原有数据
		if _, err = dao.AuthRoleAccess.Ctx(ctx).Where(do.AuthRoleAccess{
			RoleId: roleId,
		}).Delete(); err != nil {
			return err
		}
		// 写入新增数据
		if _, err = dao.AuthRoleAccess.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	// 更新授权政策
	// service.Oauth().SavePolicy(ctx)

	return nil
}

// 根据权限ID删除角色权限
func (s *sAuth) DeleteRoleAccessByRuleID(ctx context.Context, ruleId uint) error {
	var (
		err error
	)

	return dao.AuthRoleAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthRoleAccess.Ctx(ctx).Where(do.AuthRoleAccess{
			RuleId: ruleId,
		}).Delete()
		return err
	})
}

// 获取所有角色权限
func (s *sAuth) GetAllRoleAccess(ctx context.Context) ([]*entity.AuthRoleAccess, error) {
	var (
		list []*entity.AuthRoleAccess
		err  error
	)

	if err = dao.AuthRoleAccess.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}
