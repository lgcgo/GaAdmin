package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 设置分组权限
func (s *sUser) SetupRoles(ctx context.Context, userId uint, role_ids []uint) error {
	var (
		err  error
		user *entity.User
		data []*do.UserAccess
	)

	// 检测分组
	if user, err = s.GetUser(ctx, userId); err != nil {
		return err
	}
	if user == nil {
		return gerror.Newf("user is not exists: %d", userId)
	}
	// 检测角色组ID集
	if _, err = service.Auth().CheckRoleIds(ctx, role_ids); err != nil {
		return err
	}
	// 组装新增数据
	for _, v := range role_ids {
		data = append(data, &do.UserAccess{
			UserId:     userId,
			AuthRoleId: v,
		})
	}
	if err = dao.UserAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 移除原有数据
		if _, err = dao.UserAccess.Ctx(ctx).Where(do.UserAccess{
			UserId: userId,
		}).Delete(); err != nil {
			return err
		}
		// 写入新增数据
		if _, err = dao.UserAccess.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// 根据角色ID删除权限
func (s *sUser) DeleteRolesByGroupId(ctx context.Context, roleId uint) error {
	var (
		err error
	)

	return dao.UserAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserAccess.Ctx(ctx).Where(do.UserAccess{
			AuthRoleId: roleId,
		}).Delete()
		return err
	})
}

// 获取所有角色权限
func (s *sUser) GetAllRoles(ctx context.Context) ([]*entity.UserAccess, error) {
	var (
		list []*entity.UserAccess
		err  error
	)

	if err = dao.UserAccess.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}
