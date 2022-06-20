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

type cAuthRole struct{}

var AuthRole = cAuthRole{}

// 添加角色
func (c *cAuthRole) Create(ctx context.Context, req *v1.AuthRoleCreateReq) (*v1.AuthRoleCreateRes, error) {
	var (
		res    *v1.AuthRoleCreateRes
		err    error
		in     *model.AuthRoleCreateInput
		ent    *entity.AuthRole
		roleId uint
	)

	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 创建实体
	if ent, err = service.Auth().CreateRole(ctx, in); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("role is not exists: %d", roleId)
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, err
}

// 获取角色
func (c *cAuthRole) Get(ctx context.Context, req *v1.AuthRoleGetReq) (*v1.AuthRoleGetRes, error) {
	var (
		res *v1.AuthRoleGetRes
		err error
		ent *entity.AuthRole
	)

	// 获取实体
	if ent, err = service.Auth().GetRole(ctx, req.RoleId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("role not exists: %d", req.RoleId)
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 修改角色
func (c *cAuthRole) Update(ctx context.Context, req *v1.AuthRoleUpdateReq) (*v1.AuthRoleUpdateRes, error) {
	var (
		res *v1.AuthRoleUpdateRes
		err error
		in  *model.AuthRoleUpdateInput
		ent *entity.AuthRole
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if ent, err = service.Auth().UpdateRole(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 删除角色
func (c *cAuthRole) Delete(ctx context.Context, req *v1.AuthRoleDeleteReq) (*v1.AuthRoleDeleteRes, error) {
	var (
		res *v1.AuthRoleDeleteRes
		err error
	)

	// 删除实体
	if err = service.Auth().DeleteRole(ctx, req.RoleId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取角色列表
func (c *cAuthRole) Tree(ctx context.Context, req *v1.AuthRoleTreeReq) (*v1.AuthRoleTreeRes, error) {
	var (
		res *v1.AuthRoleTreeRes
		err error
		out *model.TreeDataOutput
	)

	// 获取树数据
	if out, err = service.Auth().GetRoleTreeData(ctx); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}
