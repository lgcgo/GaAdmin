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

type cOrgDepartment struct{}

var OrgDepartment = cOrgDepartment{}

// 添加分组
func (c *cOrgDepartment) Create(ctx context.Context, req *v1.OrgDepartmentCreateReq) (*v1.OrgDepartmentCreateRes, error) {
	var (
		res          *v1.OrgDepartmentCreateRes
		err          error
		in           *model.OrgDepartmentCreateInput
		ent          *entity.OrgDepartment
		departmentId uint
	)

	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 创建实体
	if ent, err = service.Org().CreateDepartment(ctx, in); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("department is not exists: %d", departmentId)
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, err
}

// 获取分组
func (c *cOrgDepartment) Get(ctx context.Context, req *v1.OrgDepartmentGetReq) (*v1.OrgDepartmentGetRes, error) {
	var (
		res *v1.OrgDepartmentGetRes
		err error
		ent *entity.OrgDepartment
	)

	// 获取实体
	if ent, err = service.Org().GetDepartment(ctx, req.DepartmentId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("department not exists: %d", req.DepartmentId)
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 修改分组
func (c *cOrgDepartment) Update(ctx context.Context, req *v1.OrgDepartmentUpdateReq) (*v1.OrgDepartmentUpdateRes, error) {
	var (
		res *v1.OrgDepartmentUpdateRes
		err error
		in  *model.OrgDepartmentUpdateInput
		ent *entity.OrgDepartment
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if ent, err = service.Org().UpdateDepartment(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 删除分组
func (c *cOrgDepartment) Delete(ctx context.Context, req *v1.OrgDepartmentDeleteReq) (*v1.OrgDepartmentDeleteRes, error) {
	var (
		res *v1.OrgDepartmentDeleteRes
		err error
	)

	// 删除实体
	if err = service.Org().DeleteMember(ctx, req.DepartmentId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取分组列表
func (c *cOrgDepartment) Tree(ctx context.Context, req *v1.OrgDepartmentTreeReq) (*v1.OrgDepartmentTreeRes, error) {
	var (
		res *v1.OrgDepartmentTreeRes
		err error
		out *model.TreeDataOutput
	)

	// 获取树数据
	if out, err = service.Org().GetDepartmentTreeData(ctx); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}
