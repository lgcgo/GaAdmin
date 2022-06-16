package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

// 根据权限ID删除分组权限
func (s *sUser) DeleteGroupAccessByRuleID(ctx context.Context, ruleId uint) error {
	var (
		err error
	)
	return dao.UserGroupAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserGroupAccess.Ctx(ctx).Where(do.UserGroupAccess{
			AuthRuleId: ruleId,
		}).Delete()
		return err
	})
}

// 获取所有分组权限
func (s *sUser) GetAllGroupAccess(ctx context.Context) ([]*entity.UserGroupAccess, error) {
	var (
		list []*entity.UserGroupAccess
		err  error
	)

	if err = dao.UserGroupAccess.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}
