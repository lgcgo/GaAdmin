package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type cAuthRule struct{}

var AuthRule = cAuthRule{}

// 添加权限节点
func (c *cAuthRule) Create(ctx context.Context, req *v1.AuthRuleCreateReq) (*v1.AuthRuleCreateRes, error) {
	var (
		ser    = service.Auth()
		res    *v1.AuthRuleCreateRes
		err    error
		in     *model.AuthRuleCreateInput
		ent    *entity.AuthRule
		ruleId uint
	)

	// 格式化创建
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if ruleId, err = ser.CreateRule(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetRule(ctx, ruleId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 获取权限节点
func (c *cAuthRule) Get(ctx context.Context, req *v1.AuthRuleGetReq) (*v1.AuthRuleGetRes, error) {
	var (
		res *v1.AuthRuleGetRes
		err error
		ent *entity.AuthRule
	)

	// 获取实体
	if ent, err = service.Auth().GetRule(ctx, req.RuleId); err != nil {
		return nil, err
	}

	if ent == nil {
		return nil, gerror.New("rule is not exists")
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 修改权限节点
func (c *cAuthRule) Update(ctx context.Context, req *v1.AuthRuleUpdateReq) (*v1.AuthRuleUpdateRes, error) {
	var (
		ser = service.Auth()
		res *v1.AuthRuleUpdateRes
		err error
		in  *model.AuthRuleUpdateInput
		ent *entity.AuthRule
	)

	// 更新实体
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if err = ser.UpdateRule(ctx, in); err != nil {
		return nil, err
	}
	// 获取实体
	if ent, err = ser.GetRule(ctx, req.RuleId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.New("menu is not exists")
	}
	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, err
}

// 删除权限菜单
func (c *cAuthRule) Delete(ctx context.Context, req *v1.AuthRuleDeleteReq) (*v1.AuthRuleDeleteRes, error) {
	var (
		res *v1.AuthRuleDeleteRes
		err error
	)

	// 删除实体
	if err = service.Auth().DeleteRule(ctx, req.RuleId); err != nil {
		return nil, err
	}

	return res, nil
}
