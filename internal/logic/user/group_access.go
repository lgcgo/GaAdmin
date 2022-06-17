package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 设置分组权限
func (s *sUser) SetupGroupAccess(ctx context.Context, groupId uint, auth_rule_ids []uint) error {
	var (
		err   error
		group *entity.UserGroup
		data  []*do.UserGroupAccess
	)

	// 检测分组
	if group, err = s.GetGroup(ctx, groupId); err != nil {
		return err
	}
	if group == nil {
		return gerror.Newf("group is not exists: %d", groupId)
	}
	// 检测权限ID集
	if _, err = service.Auth().CheckRulesIds(ctx, auth_rule_ids); err != nil {
		return err
	}
	// 组装新增数据
	for _, v := range auth_rule_ids {
		data = append(data, &do.UserGroupAccess{
			GroupId:    groupId,
			AuthRuleId: v,
		})
	}
	if err = dao.UserGroupAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 移除原有数据
		if _, err = dao.UserGroupAccess.Ctx(ctx).Where(do.UserGroupAccess{
			GroupId: groupId,
		}).Delete(); err != nil {
			return err
		}
		// 写入新增数据
		if _, err = dao.UserGroupAccess.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
}

// 根据权限ID删除分组权限
func (s *sUser) DeleteGroupAccessByRuleID(ctx context.Context, ruleId uint) error {
	var (
		err error
	)

	return dao.UserGroupAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserGroupAccess.Ctx(ctx).Where(do.UserGroupAccess{
			AuthRuleId: ruleId,
		}).Delete()
		return err
	})
}

// 获取所有分组权限
func (s *sUser) GetAllGroupAccess(ctx context.Context) ([]*entity.UserGroupAccess, error) {
	var (
		list []*entity.UserGroupAccess
		err  error
	)

	if err = dao.UserGroupAccess.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}
