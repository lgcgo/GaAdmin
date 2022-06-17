package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type cUser struct{}

var User = cUser{}

// 创建用户
func (c *cUser) Create(ctx context.Context, req *v1.UserCreateReq) (*v1.UserCreateRes, error) {
	var (
		ser = service.User()
		res *v1.UserCreateRes
		err error
		in  *model.UserCreateInput
		ent *entity.User
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 创建实体
	if ent, err = ser.CreateUser(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取用户
func (c *cUser) Get(ctx context.Context, req *v1.UserGetReq) (*v1.UserGetRes, error) {
	var (
		res *v1.UserGetRes
		err error
		ent *entity.User
	)

	// 获取实体
	if ent, err = service.User().GetUser(ctx, req.UserId); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 修改用户
func (c *cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (*v1.UserUpdateRes, error) {
	var (
		ser = service.User()
		res *v1.UserUpdateRes
		err error
		in  *model.UserUpdateInput
		ent *entity.User
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 更新实体
	if ent, err = ser.UpdateUser(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 删除用户
func (c *cUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (*v1.UserDeleteRes, error) {
	var (
		res *v1.UserDeleteRes
		err error
	)

	// 删除实体
	if err = service.User().DeleteUser(ctx, req.UserId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取员工列表
func (c *cUser) List(ctx context.Context, req *v1.UserListReq) (*v1.UserListRes, error) {
	var (
		res *v1.UserListRes
		err error
		in  *model.Page
		out *model.UserPageOutput
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 获取分页
	if out, err = service.User().GetUserPage(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}
