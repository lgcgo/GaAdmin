package org

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lgcgo/tree"
)

// 创建部门
func (s *sOrg) CreateDepartment(ctx context.Context, in *model.OrgDepartmentCreateInput) (*entity.OrgDepartment, error) {
	var (
		data     *do.OrgDepartment
		ent      *entity.OrgDepartment
		err      error
		insertId int64
	)

	// 检测父级
	if in.ParentId > 0 {
		var parent *entity.OrgDepartment
		parent, err = s.GetDepartment(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 创建实体
	if err = dao.OrgDepartment.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.OrgDepartment.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	// 获取实体
	ent, _ = s.GetDepartment(ctx, uint(insertId))

	return ent, nil
}

// 获取组织
func (s *sOrg) GetDepartment(ctx context.Context, departmentId uint) (*entity.OrgDepartment, error) {
	var (
		ent *entity.OrgDepartment
		err error
	)

	// 扫描数据
	if err = dao.OrgDepartment.Ctx(ctx).Where(do.OrgDepartment{
		Id: departmentId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改部门
func (s *sOrg) UpdateDepartment(ctx context.Context, in *model.OrgDepartmentUpdateInput) (*entity.OrgDepartment, error) {
	var (
		data *do.OrgDepartment
		ent  *entity.OrgDepartment
		err  error
	)

	// 扫描数据
	if ent, err = s.GetDepartment(ctx, in.DepartmentId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("department is not exists: %d", in.DepartmentId)
	}
	// 检测父级
	if in.ParentId > 0 {
		var (
			parent *entity.OrgDepartment
			ids    []uint
		)
		parent, err = s.GetDepartment(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
		if ids, err = s.GetDepartmentChildrenIDs(ctx, in.DepartmentId); err != nil {
			return nil, err
		}
		for _, v := range ids {
			if in.ParentId == v {
				return nil, gerror.Newf("parent can not be self or child: %d", in.ParentId)
			}
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.OrgDepartment.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.OrgDepartment.Ctx(ctx).Where(do.OrgDepartment{
			Id: in.DepartmentId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	// 获取实体
	ent, _ = s.GetDepartment(ctx, in.DepartmentId)

	return ent, nil
}

// 删除部门(硬删除)
func (s *sOrg) DeleteDepartment(ctx context.Context, departmentId uint) error {
	var (
		ent *entity.OrgDepartment
		err error
		ids []uint
	)

	// 扫描数据
	if ent, err = s.GetDepartment(ctx, departmentId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("department is not exists: %d", departmentId)
	}
	// 获取子ID集
	if ids, err = s.GetDepartmentChildrenIDs(ctx, departmentId); err != nil {
		return err
	}
	// 删除实体
	if err = dao.OrgDepartment.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.OrgDepartment.Ctx(ctx).WhereIn("id", ids).Delete()
		return err
	}); err != nil {
		return err
	}

	return nil
}

// 获取所有部门
func (s *sOrg) GetAllDepartment(ctx context.Context) ([]*entity.OrgDepartment, error) {
	var (
		list []*entity.OrgDepartment
		err  error
	)

	if err = dao.OrgDepartment.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取部门树数据
func (s *sOrg) GetDepartmentTreeData(ctx context.Context) (*model.TreeDataOutput, error) {
	var (
		t   *tree.Tree
		out *model.TreeDataOutput
		err error
	)
	if t, err = s.getDepartmentTree(ctx); err != nil {
		return nil, err
	}
	out = &model.TreeDataOutput{
		TreeData: t.GetTreeData(),
		Total:    uint(t.CountTreeData()),
	}

	return out, nil
}

// 获取菜单子ID集
func (s *sOrg) GetDepartmentChildrenIDs(ctx context.Context, departmentId uint) ([]uint, error) {
	var (
		t    *tree.Tree
		err  error
		keys []string
		ids  []uint
	)

	// 获取树对象
	if t, err = s.getDepartmentTree(ctx); err != nil {
		return nil, err
	}
	// 获取子健集
	if keys, err = t.GetSpecChildKeys(gconv.String(departmentId)); err != nil {
		return nil, err
	}
	// 格式转换
	for _, v := range keys {
		ids = append(ids, gconv.Uint(v))
	}

	return ids, nil
}

// 检测部门ID集
func (s *sOrg) CheckDepartmentIds(ctx context.Context, departmentIds []uint) ([]uint, error) {
	var (
		m    = dao.OrgDepartment.Ctx(ctx)
		err  error
		list []*entity.OrgDepartment
		res  []uint
	)

	arr := garray.NewIntArray(true)
	for _, departmentId := range departmentIds {
		arr.Append(int(departmentId))
	}
	if err = m.Fields("id").Where("id IN(?)", departmentIds).Scan(&list); err != nil {
		return nil, err
	}
	for _, v := range list {
		arr.RemoveValue(int(v.Id))
	}
	if !arr.IsEmpty() {
		arr.Iterator(func(k int, v int) bool {
			res = append(res, uint(v))
			return true
		})
		return res, gerror.Newf("department_ids is unavailable: %s", arr.String())
	}

	return nil, nil
}

// 获取部门树
func (s *sOrg) getDepartmentTree(ctx context.Context) (*tree.Tree, error) {
	var (
		list []*entity.OrgDepartment
		out  *tree.Tree
		err  error
	)

	// 获取全部数据
	if list, err = s.GetAllDepartment(ctx); err != nil {
		return nil, err
	}

	var (
		data      = make([]*tree.TreeData, 0)
		key       string
		parentKey string
	)

	for _, v := range list {
		key = gconv.String(v.Id)
		if v.ParentId > 0 {
			parentKey = gconv.String(v.ParentId)
		}
		data = append(data, &tree.TreeData{
			Title:     v.Title,
			Key:       key,
			ParentKey: parentKey,
			Value:     "",
			Weight:    v.Weigh,
		})
	}
	if out, err = tree.NewWithData(data); err != nil {
		return nil, err
	}

	return out, nil
}
