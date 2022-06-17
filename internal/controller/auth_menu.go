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

type cAuthMenu struct{}

var AuthMenu = cAuthMenu{}

// 添加菜单
func (c *cAuthMenu) Create(ctx context.Context, req *v1.AuthMenuCreateReq) (*v1.AuthMenuCreateRes, error) {
	var (
		ser = service.Auth()
		res *v1.AuthMenuCreateRes
		err error
		in  *model.AuthMenuCreateInput
		ent *entity.AuthMenu
	)

	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 创建实体
	if ent, err = ser.CreateMenu(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取菜单
func (c *cAuthMenu) Get(ctx context.Context, req *v1.AuthMenuGetReq) (*v1.AuthMenuGetRes, error) {
	var (
		res *v1.AuthMenuGetRes
		err error
		ent *entity.AuthMenu
	)

	// 获取实体
	if ent, err = service.Auth().GetMenu(ctx, req.MenuId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.New("menu is not exists")
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 修改菜单
func (c *cAuthMenu) Update(ctx context.Context, req *v1.AuthMenuUpdateReq) (*v1.AuthMenuUpdateRes, error) {
	var (
		ser = service.Auth()
		res *v1.AuthMenuUpdateRes
		err error
		in  *model.AuthMenuUpdateInput
		ent *entity.AuthMenu
	)

	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 更新实体
	if ent, err = ser.UpdateMenu(ctx, in); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(ent, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 删除菜单
func (c *cAuthMenu) Delete(ctx context.Context, req *v1.AuthMenuDeleteReq) (*v1.AuthMenuDeleteRes, error) {
	var (
		res *v1.AuthMenuDeleteRes
		err error
	)

	// 删除实体
	if err = service.Auth().DeleteMenu(ctx, req.MenuId); err != nil {
		return nil, err
	}

	return res, nil
}

// 获取菜单树
func (c *cAuthMenu) Tree(ctx context.Context, req *v1.AuthMenuTreeReq) (*v1.AuthMenuTreeRes, error) {
	var (
		res *v1.AuthMenuTreeRes
		err error
		out *model.TreeDataOutput
	)

	// 获取菜单树(全部数据)
	if out, err = service.Auth().GetMenuTreeData(ctx); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}
