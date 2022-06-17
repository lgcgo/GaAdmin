package auth

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// 添加规则
func (s *sAuth) CreateRule(ctx context.Context, in *model.AuthRuleCreateInput) (*entity.AuthRule, error) {
	var (
		available bool
		err       error
		ent       *entity.AuthRule
	)

	// 检测菜单
	if in.MenuId > 0 {
		var parent *entity.AuthMenu
		parent, err = s.GetMenu(ctx, in.MenuId)
		if parent == nil {
			return nil, gerror.Newf("menu is not exists: %d", in.MenuId)
		}
	}
	// 路径防重
	if available, err = s.IsRulePathMethodAvailable(ctx, in.Path, in.Method); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("path is already exists: %s %s", in.Path, in.Method)
	}
	// 插入数据
	var (
		data     *do.AuthRule
		insertId int64
	)
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	err = dao.AuthRule.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.AuthRule.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	})
	if ent, err = s.GetRule(ctx, uint(insertId)); err != nil {
		return nil, err
	}

	return ent, err
}

// 获取规则
func (s *sAuth) GetRule(ctx context.Context, nodeId uint) (*entity.AuthRule, error) {
	var (
		ent *entity.AuthRule
		err error
	)

	// 扫描数据
	if err = dao.AuthRule.Ctx(ctx).Where(do.AuthRule{
		Id: nodeId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改规则
func (s *sAuth) UpdateRule(ctx context.Context, in *model.AuthRuleUpdateInput) (*entity.AuthRule, error) {
	var (
		data      *do.AuthRule
		ent       *entity.AuthRule
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetRule(ctx, in.RuleId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("rule is not exists: %d", in.RuleId)
	}
	// 检测菜单
	if in.MenuId > 0 {
		var parent *entity.AuthMenu
		parent, err = s.GetMenu(ctx, in.MenuId)
		if parent == nil {
			return nil, gerror.Newf("menu is not exists: %d", in.MenuId)
		}
	}
	// 路径防重
	if available, err = s.IsRulePathMethodAvailable(ctx, in.Path, in.Method, []uint{ent.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("path is already exists: %s %s", in.Path, in.Method)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.AuthRule.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthRule.Ctx(ctx).Where(do.AuthRule{
			Id: in.RuleId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	ent, _ = s.GetRule(ctx, in.RuleId)
	return ent, nil
}

// 删除规则(硬删除)
func (s *sAuth) DeleteRule(ctx context.Context, ruleId uint) error {
	var (
		ent *entity.AuthRule
		err error
	)

	// 扫描数据
	if ent, err = s.GetRule(ctx, ruleId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("node is not exists: %d", ruleId)
	}

	// 删除数据
	if err = dao.AuthRule.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, err = dao.AuthRule.Ctx(ctx).Where(do.AuthRule{
			Id: ruleId,
		}).Delete(); err != nil {
			return err
		}
		if err = service.User().DeleteGroupAccessByRuleID(ctx, ruleId); err != nil {
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

// 检测规则路径&请求方法
func (s *sAuth) IsRulePathMethodAvailable(ctx context.Context, path string, method string, notIds ...uint) (bool, error) {
	var (
		m     = dao.AuthRule.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.AuthRule.Columns().Id, v)
	}
	if count, err = m.Where(do.AuthRule{
		Path:   path,
		Method: method,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 检测权限ID集
func (s *sAuth) CheckRulesIds(ctx context.Context, ruleIds []uint) ([]uint, error) {
	var (
		m    = dao.AuthRule.Ctx(ctx)
		err  error
		list []*entity.AuthRule
		res  []uint
	)

	arr := garray.NewIntArray(true)
	for _, ruleId := range ruleIds {
		arr.Append(int(ruleId))
	}
	if err = m.Fields("id").Where("id IN(?)", ruleIds).Scan(&list); err != nil {
		return nil, err
	}
	for _, v := range list {
		arr.RemoveValue(int(v.Id))
	}
	if !arr.IsEmpty() {
		arr.Iterator(func(k int, v int) bool {
			res = append(res, uint(v))
			return true
		})
		return res, gerror.Newf("rule_ids is unavailable: %s", arr.String())
	}

	return nil, nil
}
