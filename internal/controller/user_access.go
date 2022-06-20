package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/service"
	"context"
)

type cUserAccess struct{}

var UserAccess = cUserAccess{}

func (c *cUserAccess) Setup(ctx context.Context, req *v1.UserAccessSetupReq) (*v1.UserAccessSetupRes, error) {
	var (
		res *v1.UserAccessSetupRes
		err error
	)

	// 设置用户组权限
	if err = service.User().SetupRoles(ctx, req.UserId, req.RoleIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.UserAccessSetupRes{
		UserId:  req.UserId,
		RoleIds: req.RoleIds,
	}

	return res, nil
}
