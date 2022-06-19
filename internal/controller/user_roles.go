package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/service"
	"context"
)

type cUserRoles struct{}

var UserRoles = cUserRoles{}

func (c *cUserRoles) Setup(ctx context.Context, req *v1.UserRolesSetupReq) (*v1.UserRolesSetupRes, error) {
	var (
		res *v1.UserRolesSetupRes
		err error
	)

	// 设置用户组权限
	if err = service.User().SetupRoles(ctx, req.UserId, req.GroupIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.UserRolesSetupRes{
		UserId:   req.UserId,
		GroupIds: req.GroupIds,
	}

	return res, nil
}
