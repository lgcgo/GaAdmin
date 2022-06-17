package org

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOrg struct{}

func init() {
	service.RegisterOrg(New())
}

func New() *sOrg {
	return &sOrg{}
}

// 创建组织
func (s *sOrg) CreateOrg(ctx context.Context, in *model.OrgCreateInput) (*entity.Org, error) {
	var (
		available bool
		err       error
		ent       *entity.Org
	)

	// 验证组织名称
	if available, err = s.IsNameAvailable(ctx, in.Name); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("name is already exists: %s", in.Name)
	}

	// 验证组织编码，如果有
	if len(in.CertificatesNo) > 0 {
		if available, err = s.IsCertificatesNoAvailable(ctx, in.CertificatesNo); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("certificatesNo is already exists: %s", in.CertificatesNo)
		}
	}

	// 插入数据
	var (
		data     *do.Org
		insertId int64
	)

	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if err = dao.Org.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.Org.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	if ent, err = s.GetOrg(ctx, uint(insertId)); err != nil {
		return nil, err
	}

	return ent, nil
}

// 获取组织
func (s *sOrg) GetOrg(ctx context.Context, orgId uint) (*entity.Org, error) {
	var (
		ent *entity.Org
		err error
	)

	// 扫描数据
	if err = dao.Org.Ctx(ctx).Where(do.Org{
		Id: orgId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改组织
func (s *sOrg) UpdateOrg(ctx context.Context, in *model.OrgUpdateInput) (*entity.Org, error) {
	var (
		data      *do.Org
		ent       *entity.Org
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetOrg(ctx, in.OrgId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("org is not exist: %d", in.OrgId)
	}

	// 验证组织名称
	if available, err = s.IsNameAvailable(ctx, in.Name, []uint{ent.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("name is already exists: %s", in.Name)
	}

	// 验证组织编码，如果有
	if len(in.CertificatesNo) > 0 {
		if available, err = s.IsCertificatesNoAvailable(ctx, in.CertificatesNo, []uint{ent.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("certificatesNo is already exists: %s", in.CertificatesNo)
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.Org.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().Id, in.OrgId).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	ent, _ = s.GetOrg(ctx, in.OrgId)

	return ent, nil
}

// 删除组织
func (s *sOrg) DeleteOrg(ctx context.Context, id uint) error {
	var (
		ent *entity.Org
		err error
	)

	// 扫描数据
	if ent, err = s.GetOrg(ctx, id); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("org is not exists: %d", id)
	}

	// 删除数据
	return dao.Org.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Org.Ctx(ctx).Where(dao.Org.Columns().Id, id).Delete()
		return err
	})
}

// 组织分页
func (s *sOrg) GetOrgPage(ctx context.Context, in *model.Page) (*model.OrgPageOutput, error) {
	var (
		m    = dao.Org.Ctx(ctx)
		out  = &model.OrgPageOutput{}
		list []*entity.Org
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

// 检测组织名称
func (s *sOrg) IsNameAvailable(ctx context.Context, name string, notIds ...uint) (bool, error) {
	var (
		m     = dao.Org.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.Org.Columns().Id, v)
	}
	if count, err = m.Where(g.Map{
		dao.Org.Columns().Name: name,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 检查组织证件号码
func (s *sOrg) IsCertificatesNoAvailable(ctx context.Context, certificatesNo string, notIds ...uint) (bool, error) {
	var (
		m     = dao.Org.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.Org.Columns().Id, v)
	}
	if count, err = m.Where(g.Map{
		dao.Org.Columns().CertificatesNo: certificatesNo,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}
