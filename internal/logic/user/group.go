package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lgcgo/tree"
)

// 创建分组
func (s *sUser) CreateGroup(ctx context.Context, in *model.UserGroupCreateInput) (*entity.UserGroup, error) {
	var (
		data      *do.UserGroup
		ent       *entity.UserGroup
		err       error
		available bool
		insertId  int64
	)

	// 检测父级
	if in.ParentId > 0 {
		var parent *entity.UserGroup
		parent, err = s.GetGroup(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
	}
	// 路径防重
	if available, err = s.isGroupNameAvailable(ctx, in.Name); err != nil {
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
	if err = dao.UserGroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.UserGroup.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)
	// 获取实体
	ent, _ = s.GetGroup(ctx, uint(insertId))

	return ent, nil
}

// 获取组织
func (s *sUser) GetGroup(ctx context.Context, groupId uint) (*entity.UserGroup, error) {
	var (
		ent *entity.UserGroup
		err error
	)

	// 扫描数据
	if err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
		Id: groupId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改分组
func (s *sUser) UpdateGroup(ctx context.Context, in *model.UserGroupUpdateInput) (*entity.UserGroup, error) {
	var (
		data      *do.UserGroup
		ent       *entity.UserGroup
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetGroup(ctx, in.GroupId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("group is not exists: %d", in.GroupId)
	}
	// 检测父级
	if in.ParentId > 0 {
		var (
			parent *entity.UserGroup
			ids    []uint
		)
		parent, err = s.GetGroup(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
		if ids, err = s.GetGroupChildrenIDs(ctx, in.GroupId); err != nil {
			return nil, err
		}
		for _, v := range ids {
			if in.ParentId == v {
				return nil, gerror.Newf("parent can not be self or child: %d", in.ParentId)
			}
		}
	}
	// 名称防重
	if available, err = s.isGroupNameAvailable(ctx, in.Name, []uint{ent.Id}...); err != nil {
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
	if err = dao.UserGroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
			Id: in.GroupId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)
	// 获取实体
	ent, _ = s.GetGroup(ctx, in.GroupId)

	return ent, nil
}

// 删除分组(硬删除)
func (s *sUser) DeleteGroup(ctx context.Context, groupId uint) error {
	var (
		ent *entity.UserGroup
		err error
		ids []uint
	)

	// 扫描数据
	if ent, err = s.GetGroup(ctx, groupId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("group is not exists: %d", groupId)
	}
	// 获取子ID集
	if ids, err = s.GetGroupChildrenIDs(ctx, groupId); err != nil {
		return err
	}
	// 删除实体
	if err = dao.UserGroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserGroup.Ctx(ctx).WhereIn("id", ids).Delete()
		return err
	}); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
}

// 获取所有分组
func (s *sUser) GetAllGroup(ctx context.Context) ([]*entity.UserGroup, error) {
	var (
		list []*entity.UserGroup
		err  error
	)

	if err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
		Status: "normal",
	}).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取分组树数据
func (s *sUser) GetGroupTreeData(ctx context.Context) (*model.TreeDataOutput, error) {
	var (
		t   *tree.Tree
		out *model.TreeDataOutput
		err error
	)
	if t, err = s.getGroupTree(ctx); err != nil {
		return nil, err
	}
	out = &model.TreeDataOutput{
		TreeData: t.GetTreeData(),
		Total:    uint(t.CountTreeData()),
	}

	return out, nil
}

// 获取分组名称
func (s *sUser) GetGroupName(ctx context.Context, gourpID uint) (string, error) {
	var (
		err error
		val *gvar.Var
	)

	if val, err = dao.UserGroup.Ctx(ctx).Fields("name").Where(do.UserGroup{
		Id: gourpID,
	}).Value(); err != nil {
		return "", err
	}

	return val.String(), nil
}

// 获取菜单子ID集
func (s *sUser) GetGroupChildrenIDs(ctx context.Context, groupId uint) ([]uint, error) {
	var (
		t    *tree.Tree
		err  error
		keys []string
		ids  []uint
	)

	// 获取树对象
	if t, err = s.getGroupTree(ctx); err != nil {
		return nil, err
	}
	// 获取子健集
	if keys, err = t.GetSpecChildKeys(gconv.String(groupId)); err != nil {
		return nil, err
	}
	// 格式转换
	for _, v := range keys {
		ids = append(ids, gconv.Uint(v))
	}

	return ids, nil
}

// 检测分组ID集
func (s *sUser) CheckGroupIds(ctx context.Context, groupIds []uint) ([]uint, error) {
	var (
		m    = dao.UserGroup.Ctx(ctx)
		err  error
		list []*entity.UserGroup
		res  []uint
	)

	arr := garray.NewIntArray(true)
	for _, groupId := range groupIds {
		arr.Append(int(groupId))
	}
	if err = m.Fields("id").Where("id IN(?)", groupIds).Scan(&list); err != nil {
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
		return res, gerror.Newf("group_ids is unavailable: %s", arr.String())
	}

	return nil, nil
}

// 获取分组树
func (s *sUser) getGroupTree(ctx context.Context) (*tree.Tree, error) {
	var (
		list []*entity.UserGroup
		out  *tree.Tree
		err  error
	)

	// 获取全部数据
	if list, err = s.GetAllGroup(ctx); err != nil {
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

// 检测分组名称
func (s *sUser) isGroupNameAvailable(ctx context.Context, name string, notIds ...uint) (bool, error) {
	var (
		m     = dao.UserGroup.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.UserGroup.Columns().Id, v)
	}
	if count, err = m.Where(do.UserGroup{
		Name: name,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}
