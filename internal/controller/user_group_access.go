package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/service"
	"context"
)

type cUserGroupAccess struct{}

var UserGroupAccess = cUserGroupAccess{}

func (c *cUserGroupAccess) Setup(ctx context.Context, req *v1.UserGroupAccessSetupReq) (*v1.UserGroupAccessSetupRes, error) {
	var (
		res *v1.UserGroupAccessSetupRes
		err error
	)

	// 设置用户组权限
	if err = service.User().SetupGroupAccess(ctx, req.GroupId, req.AuthRuleIds); err != nil {
		return nil, err
	}
	// 转换响应
	res = &v1.UserGroupAccessSetupRes{
		GroupId:     req.GroupId,
		AuthRuleIds: req.AuthRuleIds,
	}

	return res, nil
}
