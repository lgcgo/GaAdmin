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

type cOrg struct{}

var Org = &cOrg{}

// 创建
func (c *cOrg) Create(ctx context.Context, req *v1.OrgCreateReq) (*v1.OrgCreateRes, error) {
	var (
		ser   = service.Org()
		res   *v1.OrgCreateRes
		err   error
		in    *model.OrgCreateInput
		ent   *entity.Org
		orgId uint
	)

	// 格式化创建
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if orgId, err = ser.CreateOrg(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetOrg(ctx, orgId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 获取
func (c *cOrg) Get(ctx context.Context, req *v1.OrgGetReq) (*v1.OrgGetRes, error) {
	var (
		res *v1.OrgGetRes
		err error
		ent *entity.Org
	)

	// 获取实体
	if ent, err = service.Org().GetOrg(ctx, req.OrgId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("org is not exist: %d", req.OrgId)
	}
	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 修改
func (c *cOrg) Update(ctx context.Context, req *v1.OrgUpdateReq) (*v1.OrgUpdateRes, error) {
	var (
		ser = service.Org()
		res *v1.OrgUpdateRes
		err error
		in  *model.OrgUpdateInput
		ent *entity.Org
	)

	// 格式化更新
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if err = ser.UpdateOrg(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetOrg(ctx, req.OrgId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 删除
func (c *cOrg) Delete(ctx context.Context, req *v1.OrgDeleteReq) (*v1.OrgDeleteRes, error) {
	var (
		res *v1.OrgDeleteRes
		err error
	)

	// 删除实体
	err = service.Org().DeleteOrg(ctx, req.OrgId)

	return res, err
}

// 列表
func (c *cOrg) List(ctx context.Context, req *v1.OrgListReq) (*v1.OrgListRes, error) {
	var (
		res *v1.OrgListRes
		err error
		in  *model.Page
		out *model.OrgPageOutput
	)

	// 格式化获取分页
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if out, err = service.Org().GetOrgPage(ctx, in); err != nil {
		return nil, err
	}

	// 格式化返回
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}
	return res, err
}
