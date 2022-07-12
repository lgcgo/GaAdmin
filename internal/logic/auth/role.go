package auth

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lgcgo/tree"
)

// 创建角色
func (s *sAuth) CreateRole(ctx context.Context, in *model.AuthRoleCreateInput) (*entity.AuthRole, error) {
	var (
		data      *do.AuthRole
		ent       *entity.AuthRole
		err       error
		available bool
		insertId  int64
	)

	// 检测父级
	if in.ParentId > 0 {
		var parent *entity.AuthRole
		parent, err = s.GetRole(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
	}
	// 名称防重
	if available, err = s.isRoleNameAvailable(ctx, in.Name); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("name is already exists: %s", in.Name)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 创建实体
	if err = dao.AuthRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.AuthRole.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	// 更新授权政策
	// service.Oauth().SavePolicy(ctx)
	// 获取实体
	ent, _ = s.GetRole(ctx, uint(insertId))

	return ent, nil
}

// 获取组织
func (s *sAuth) GetRole(ctx context.Context, roleId uint) (*entity.AuthRole, error) {
	var (
		ent *entity.AuthRole
		err error
	)

	// 扫描数据
	if err = dao.AuthRole.Ctx(ctx).Where(do.AuthRole{
		Id: roleId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改角色
func (s *sAuth) UpdateRole(ctx context.Context, in *model.AuthRoleUpdateInput) (*entity.AuthRole, error) {
	var (
		data      *do.AuthRole
		ent       *entity.AuthRole
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetRole(ctx, in.RoleId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("role is not exists: %d", in.RoleId)
	}
	// 检测父级
	if in.ParentId > 0 {
		var (
			parent *entity.AuthRole
			ids    []uint
		)
		parent, err = s.GetRole(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
		if ids, err = s.GetRoleChildrenIDs(ctx, in.RoleId); err != nil {
			return nil, err
		}
		for _, v := range ids {
			if in.ParentId == v {
				return nil, gerror.Newf("parent can not be self or child: %d", in.ParentId)
			}
		}
	}
	// 名称防重
	if available, err = s.isRoleNameAvailable(ctx, in.Name, []uint{ent.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("name is already exists: %s", in.Name)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.AuthRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthRole.Ctx(ctx).Where(do.AuthRole{
			Id: in.RoleId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	// 更新授权政策
	// service.Oauth().SavePolicy(ctx)
	// 获取实体
	ent, _ = s.GetRole(ctx, in.RoleId)

	return ent, nil
}

// 删除角色(硬删除)
func (s *sAuth) DeleteRole(ctx context.Context, roleId uint) error {
	var (
		ent *entity.AuthRole
		err error
		ids []uint
	)

	// 扫描数据
	if ent, err = s.GetRole(ctx, roleId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("role is not exists: %d", roleId)
	}
	// 获取子ID集
	if ids, err = s.GetRoleChildrenIDs(ctx, roleId); err != nil {
		return err
	}
	// 删除实体
	if err = dao.AuthRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthRole.Ctx(ctx).WhereIn("id", ids).Delete()
		return err
	}); err != nil {
		return err
	}
	// 更新授权政策
	// service.Oauth().SavePolicy(ctx)

	return nil
}

// 获取所有角色
func (s *sAuth) GetAllRole(ctx context.Context) ([]*entity.AuthRole, error) {
	var (
		list []*entity.AuthRole
		err  error
	)

	if err = dao.AuthRole.Ctx(ctx).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取角色树数据
func (s *sAuth) GetRoleTreeData(ctx context.Context) (*model.TreeDataOutput, error) {
	var (
		t   *tree.Tree
		out *model.TreeDataOutput
		err error
	)
	if t, err = s.getRoleTree(ctx); err != nil {
		return nil, err
	}
	out = &model.TreeDataOutput{
		TreeData: t.GetTreeData(),
		Total:    uint(t.CountTreeData()),
	}

	return out, nil
}

// 获取角色名称
func (s *sAuth) GetRoleName(ctx context.Context, roleID uint) (string, error) {
	var (
		err error
		val *gvar.Var
	)

	if val, err = dao.AuthRole.Ctx(ctx).Fields("name").Where(do.AuthRole{
		Id: roleID,
	}).Value(); err != nil {
		return "", err
	}

	return val.String(), nil
}

// 获取菜单子ID集
func (s *sAuth) GetRoleChildrenIDs(ctx context.Context, roleId uint) ([]uint, error) {
	var (
		t    *tree.Tree
		err  error
		keys []string
		ids  []uint
	)

	// 获取树对象
	if t, err = s.getRoleTree(ctx); err != nil {
		return nil, err
	}
	// 获取子健集
	if keys, err = t.GetSpecChildKeys(gconv.String(roleId)); err != nil {
		return nil, err
	}
	// 格式转换
	for _, v := range keys {
		ids = append(ids, gconv.Uint(v))
	}

	return ids, nil
}

// 检测角色ID集
func (s *sAuth) CheckRoleIds(ctx context.Context, roleIds []uint) ([]uint, error) {
	var (
		m    = dao.AuthRole.Ctx(ctx)
		err  error
		list []*entity.AuthRole
		res  []uint
	)

	arr := garray.NewIntArray(true)
	for _, roleId := range roleIds {
		arr.Append(int(roleId))
	}
	if err = m.Fields("id").Where("id IN(?)", roleIds).Scan(&list); err != nil {
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
		return res, gerror.Newf("role_ids is unavailable: %s", arr.String())
	}

	return nil, nil
}

// 获取角色树
func (s *sAuth) getRoleTree(ctx context.Context) (*tree.Tree, error) {
	var (
		data      = make([]*tree.TreeData, 0)
		list      []*entity.AuthRole
		out       *tree.Tree
		err       error
		key       string
		parentKey string
	)

	// 获取全部数据
	if list, err = s.GetAllRole(ctx); err != nil {
		return nil, err
	}
	// 组装树数据源
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
	// 实例化树
	if out, err = tree.NewWithData(data); err != nil {
		return nil, err
	}

	return out, nil
}

// 检测角色名称
func (s *sAuth) isRoleNameAvailable(ctx context.Context, name string, notIds ...uint) (bool, error) {
	var (
		m     = dao.AuthRole.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.AuthRole.Columns().Id, v)
	}
	if count, err = m.Where(do.AuthRole{
		Name: name,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}
