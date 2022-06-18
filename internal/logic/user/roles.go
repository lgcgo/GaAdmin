package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 设置分组权限
func (s *sUser) SetupRoles(ctx context.Context, userId uint, group_ids []uint) error {
	var (
		err  error
		user *entity.User
		data []*do.UserRoles
	)

	// 检测分组
	if user, err = s.GetUser(ctx, userId); err != nil {
		return err
	}
	if user == nil {
		return gerror.Newf("user is not exists: %d", userId)
	}
	// 检测用户组ID集
	if _, err = s.CheckGroupIds(ctx, group_ids); err != nil {
		return err
	}
	// 组装新增数据
	for _, v := range group_ids {
		data = append(data, &do.UserRoles{
			UserId:  userId,
			GroupId: v,
		})
	}
	if err = dao.UserRoles.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 移除原有数据
		if _, err = dao.UserRoles.Ctx(ctx).Where(do.UserRoles{
			UserId: userId,
		}).Delete(); err != nil {
			return err
		}
		// 写入新增数据
		if _, err = dao.UserRoles.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// 根据分组ID删除分组权限
func (s *sUser) DeleteRolesByGroupId(ctx context.Context, groupId uint) error {
	var (
		err error
	)

	return dao.UserRoles.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserRoles.Ctx(ctx).Where(do.UserRoles{
			GroupId: groupId,
		}).Delete()
		return err
	})
}

// 获取所有分组权限
func (s *sUser) GetAllRoles(ctx context.Context) ([]*entity.UserRoles, error) {
	var (
		list []*entity.UserRoles
		err  error
	)

	if err = dao.UserRoles.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}
