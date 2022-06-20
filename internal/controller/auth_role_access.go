package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/service"
	"context"
)

type cAuthRoleAccess struct{}

var AuthRoleAccess = cAuthRoleAccess{}

// 设置角色权限
func (c *cAuthRoleAccess) Setup(ctx context.Context, req *v1.AuthRoleAccessSetupReq) (*v1.AuthRoleAccessSetupRes, error) {
	var (
		res *v1.AuthRoleAccessSetupRes
		err error
	)

	// 设置用户组权限
	if err = service.Auth().SetupRoleAccess(ctx, req.RoleId, req.AuthRuleIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.AuthRoleAccessSetupRes{
		RoleId:      req.RoleId,
		AuthRuleIds: req.AuthRuleIds,
	}

	return res, nil
}
