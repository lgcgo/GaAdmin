package org

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// 创建成员
func (s *sOrg) CreateMember(ctx context.Context, in *model.OrgMemberCreateInput) (*entity.OrgMember, error) {
	var (
		available bool
		err       error
		ent       *entity.OrgMember
	)

	// 检测会员ID
	if available, err = s.isMemberUserIdAvailable(ctx, in.OrgId, in.UserId); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("user is already exists: %s", in.UserId)
	}

	// 验证成员编号，如果有
	if len(in.No) > 0 {
		if available, err = s.isMemberNoAvailable(ctx, in.OrgId, in.No); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("no is already exists: %s", in.No)
		}
	}

	// 插入数据
	var (
		data     *entity.OrgMember
		insertId int64
	)
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	data.Status = "normal"
	if err = dao.OrgMember.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.OrgMember.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	if ent, err = s.GetMember(ctx, uint(insertId)); err != nil {
		return nil, err
	}

	return ent, nil
}

// 获取成员
func (s *sOrg) GetMember(ctx context.Context, memberId uint) (*entity.OrgMember, error) {
	var (
		ent *entity.OrgMember
		err error
	)

	// 扫描数据
	if err = dao.OrgMember.Ctx(ctx).Where(do.OrgMember{
		Id: memberId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改成员
func (s *sOrg) UpdateMember(ctx context.Context, in *model.OrgMemberUpdateInput) (*entity.OrgMember, error) {
	var (
		data      *do.OrgMember
		ent       *entity.OrgMember
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetMember(ctx, in.MemberId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("member is no exists: %d", in.MemberId)
	}
	// 验证成员编号，如果有
	if len(in.No) > 0 {
		if available, err = s.isMemberNoAvailable(ctx, ent.OrgId, in.No, []uint{ent.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("no is already exists: %s", in.No)
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.OrgMember.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.OrgMember.Ctx(ctx).Where(dao.OrgMember.Columns().Id, in.MemberId).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	ent, _ = s.GetMember(ctx, in.MemberId)

	return ent, nil
}

// 删除成员
func (s *sOrg) DeleteMember(ctx context.Context, id uint) error {
	var (
		err error
		ent *entity.OrgMember
	)

	// 扫描数据
	if ent, err = s.GetMember(ctx, id); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("org is not exists: %d", id)
	}

	// 删除数据
	return dao.OrgMember.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.OrgMember.Ctx(ctx).Where(dao.OrgMember.Columns().Id, id).Delete()
		return err
	})
}

// 成员分页
func (s *sOrg) GetMemberPage(ctx context.Context, in *model.Page) (*model.OrgMemberPageOutput, error) {
	var (
		m    = dao.OrgMember.Ctx(ctx)
		out  = &model.OrgMemberPageOutput{}
		list []*entity.OrgMember
		err  error
	)

	// 分页默认值
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}

	// 组装条件
	if len(in.Condition) > 0 {
		m.Where(in.Condition)
	}

	// 扫描列表
	if err = m.Page(in.Page, in.Size).Order(in.Order).Scan(&list); err != nil {
		return nil, err
	}
	out.List = list

	// 统计分页
	if out.Pager.Total, err = m.Count(); err != nil {
		return nil, err
	}
	out.Size = in.Size
	out.Page = in.Page

	return out, err
}

// 检测成员编号
func (s *sOrg) isMemberUserIdAvailable(ctx context.Context, orgId uint, userId uint, notIds ...uint) (bool, error) {
	var (
		m     = dao.OrgMember.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.OrgMember.Columns().Id, v)
	}
	if count, err = m.Where(do.OrgMember{
		OrgId:  orgId,
		UserId: userId,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 检测成员编号
func (s *sOrg) isMemberNoAvailable(ctx context.Context, orgId uint, no string, notIds ...uint) (bool, error) {
	var (
		m     = dao.OrgMember.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.OrgMember.Columns().Id, v)
	}
	if count, err = m.Where(do.OrgMember{
		OrgId: orgId,
		No:    no,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}
