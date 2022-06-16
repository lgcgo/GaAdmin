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

type cUserGroup struct{}

var UserGroup = cUserGroup{}

// 添加分组
func (c *cUserGroup) Create(ctx context.Context, req *v1.UserGroupCreateReq) (*v1.UserGroupCreateRes, error) {
	var (
		ser     = service.User()
		res     *v1.UserGroupCreateRes
		err     error
		in      *model.UserGroupCreateInput
		ent     *entity.UserGroup
		groupId uint
	)

	// 格式化创建
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if groupId, err = ser.CreateGroup(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetGroup(ctx, groupId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("group is not exists: %d", groupId)
	}
	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 获取分组
func (c *cUserGroup) Get(ctx context.Context, req *v1.UserGroupGetReq) (*v1.UserGroupGetRes, error) {
	var (
		res *v1.UserGroupGetRes
		err error
		ent *entity.UserGroup
	)

	// 获取实体
	if ent, err = service.User().GetGroup(ctx, req.GroupId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("group not exists: %d", req.GroupId)
	}
	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 修改分组
func (c *cUserGroup) Update(ctx context.Context, req *v1.UserGroupUpdateReq) (*v1.UserGroupUpdateRes, error) {
	var (
		ser = service.User()
		res *v1.UserGroupUpdateRes
		err error
		in  *model.UserGroupUpdateInput
		ent *entity.UserGroup
	)

	// 格式化更新
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	if err = ser.UpdateGroup(ctx, in); err != nil {
		return nil, err
	}

	// 获取实体
	if ent, err = ser.GetGroup(ctx, req.GroupId); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}
	return res, err
}

// 删除分组
func (c *cUserGroup) Delete(ctx context.Context, req *v1.UserGroupDeleteReq) (*v1.UserGroupDeleteRes, error) {
	var (
		res *v1.UserGroupDeleteRes
		err error
	)

	// 删除实体
	if err = service.User().DeleteGroup(ctx, req.GroupId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取分组列表
func (c *cUserGroup) Tree(ctx context.Context, req *v1.UserGroupTreeReq) (*v1.UserGroupTreeRes, error) {
	var (
		res *v1.UserGroupTreeRes
		err error
		out *model.TreeDataOutput
	)

	// 获取分页(全部数据)
	if out, err = service.User().GetGroupTreeData(ctx); err != nil {
		return nil, err
	}

	// 格式化响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}
