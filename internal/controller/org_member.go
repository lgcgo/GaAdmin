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

type cOrgMember struct{}

var OrgMember = cOrgMember{}

// 创建成员
func (c *cOrgMember) Create(ctx context.Context, req *v1.OrgMemberCreateReq) (*v1.OrgMemberCreateRes, error) {
	var (
		ser      = service.Org()
		res      *v1.OrgMemberCreateRes
		err      error
		in       *model.OrgMemberCreateInput
		memberId uint
		ent      *entity.OrgMember
	)

	// 格式化创建
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if memberId, err = ser.CreateMember(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetMember(ctx, memberId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 获取成员
func (c *cOrgMember) Get(ctx context.Context, req *v1.OrgMemberGetReq) (*v1.OrgMemberGetRes, error) {
	var (
		res *v1.OrgMemberGetRes
		err error
		ent *entity.OrgMember
	)

	// 获取实体
	if ent, err = service.Org().GetMember(ctx, req.MemberId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("member is not exist: %d", req.MemberId)
	}
	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 修改
func (c *cOrgMember) Update(ctx context.Context, req *v1.OrgMemberUpdateReq) (*v1.OrgMemberUpdateRes, error) {
	var (
		ser = service.Org()
		res *v1.OrgMemberUpdateRes
		err error
		in  *model.OrgMemberUpdateInput
		ent *entity.OrgMember
	)

	// 格式化更新
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if err = ser.UpdateMember(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetMember(ctx, req.MemberId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 删除
func (c *cOrgMember) Delete(ctx context.Context, req *v1.OrgMemberDeleteReq) (*v1.OrgMemberDeleteRes, error) {
	var (
		res *v1.OrgMemberDeleteRes
		err error
	)

	// 删除实体
	if err = service.Org().DeleteMember(ctx, req.MemberId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取成员列表
func (c *cOrgMember) List(ctx context.Context, req *v1.OrgMemberListReq) (*v1.OrgMemberListRes, error) {
	var (
		res *v1.OrgMemberListRes
		err error
		in  *model.Page
		out *model.OrgMemberPageOutput
	)

	// 格式化获取分页
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if out, err = service.Org().GetMemberPage(ctx, in); err != nil {
		return nil, err
	}

	// 格式化返回
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}
	return res, nil
}
