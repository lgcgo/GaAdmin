package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/service"
	"context"
)

type cAuthAccess struct{}

var AuthAccess = cAuthAccess{}

func (c *cAuthAccess) SetupBasic(ctx context.Context, req *v1.AuthAccessSetupBasicReq) (*v1.AuthAccessSetupBasicRes, error) {
	var (
		res *v1.AuthAccessSetupBasicRes
		err error
	)

	// 设置基础用户权限
	if err = service.Auth().SetupBasicAccess(ctx, req.RuleIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.AuthAccessSetupBasicRes{
		RuleIds: req.RuleIds,
	}

	return res, nil
}

func (c *cAuthAccess) SetupLimited(ctx context.Context, req *v1.AuthAccessSetupLimitedReq) (*v1.AuthAccessSetupLimitedRes, error) {
	var (
		res *v1.AuthAccessSetupLimitedRes
		err error
	)

	// 设置基础用户权限
	if err = service.Auth().SetupBasicAccess(ctx, req.RuleIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.AuthAccessSetupLimitedRes{
		RuleIds: req.RuleIds,
	}

	return res, nil
}
