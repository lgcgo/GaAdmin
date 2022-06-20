package auth

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

// 获取所有基础权限
func (s *sAuth) GetBasicAccessRuleIds(ctx context.Context) ([]uint, error) {
	var (
		list []*entity.AuthAccess
		err  error
		res  []uint
	)

	if err = dao.AuthAccess.Ctx(ctx).Where(do.AuthAccess{
		Type: "basic",
	}).Scan(&list); err != nil {
		return nil, err
	}
	for _, v := range list {
		res = append(res, v.RuleId)
	}

	return res, nil
}

// 获取所有受限权限
func (s *sAuth) GetAllLimitedAccess(ctx context.Context) ([]uint, error) {
	var (
		list []*entity.AuthAccess
		err  error
		res  []uint
	)

	if err = dao.AuthAccess.Ctx(ctx).Where(do.AuthAccess{
		Type: "limited",
	}).Scan(&list); err != nil {
		return nil, err
	}
	for _, v := range list {
		res = append(res, v.RuleId)
	}

	return res, nil
}

// 设置基础权限
func (s *sAuth) SetupBasicAccess(ctx context.Context, ruleIds []uint) error {
	var (
		err error
	)

	// 设置权限
	if err = s.setupAccessAtType(ctx, "basic", ruleIds); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
}

// 设置受限权限
func (s *sAuth) SetupLimitedAccess(ctx context.Context, ruleIds []uint) error {
	var (
		err error
	)

	// 设置权限
	if err = s.setupAccessAtType(ctx, "limited", ruleIds); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
}

// 设置权限
func (s *sAuth) setupAccessAtType(ctx context.Context, typeStr string, ruleIds []uint) error {
	var (
		err  error
		data []*do.AuthAccess
	)

	// 检测权限ID集
	if err = s.CheckRuleIds(ctx, ruleIds); err != nil {
		return err
	}
	// 组装新增数据
	for _, v := range ruleIds {
		data = append(data, &do.AuthAccess{
			Type:   "limited",
			RuleId: v,
		})
	}

	return dao.AuthAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 移除原有数据
		if _, err = dao.AuthAccess.Ctx(ctx).Delete(); err != nil {
			return err
		}
		// 写入新增数据
		if _, err = dao.AuthAccess.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	})
}
