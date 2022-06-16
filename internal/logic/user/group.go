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
func (s *sUser) CreateGroup(ctx context.Context, in *model.UserGroupCreateInput) (uint, error) {
	var (
		available bool
		err       error
	)

	// 检测父级
	if in.ParentId > 0 {
		var parent *entity.UserGroup
		parent, err = s.GetGroup(ctx, in.ParentId)
		if parent == nil {
			return 0, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
	}
	// 路径防重
	if available, err = s.IsGroupNameAvailable(ctx, in.Name); err != nil {
		return 0, err
	}
	if !available {
		return 0, gerror.Newf("name is already exists: %s", in.Name)
	}
	// 插入数据
	var (
		data     *do.UserGroup
		insertId int64
	)
	if err = gconv.Struct(in, &data); err != nil {
		return 0, err
	}
	if err = dao.UserGroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.UserGroup.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return 0, err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return uint(insertId), nil
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
func (s *sUser) UpdateGroup(ctx context.Context, in *model.UserGroupUpdateInput) error {
	var (
		ent       *entity.UserGroup
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetGroup(ctx, in.GroupId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("group is not exists: %d", in.GroupId)
	}
	// 检测父级
	if in.ParentId > 0 {
		var (
			parent *entity.UserGroup
			ids    []uint
		)
		parent, err = s.GetGroup(ctx, in.ParentId)
		if parent == nil {
			return gerror.Newf("parent is not exists: %d", in.ParentId)
		}
		if ids, err = s.GetGroupChildrenIDs(ctx, in.GroupId); err != nil {
			return err
		}
		for _, v := range ids {
			if in.ParentId == v {
				return gerror.Newf("parent can not be self or child: %d", in.ParentId)
			}
		}
	}
	// 名称防重
	if available, err = s.IsGroupNameAvailable(ctx, in.Name, []uint{ent.Id}...); err != nil {
		return err
	}
	if !available {
		return gerror.Newf("name is already exists: %s", in.Name)
	}
	// 格式化更新
	var data *do.UserGroup
	if err = gconv.Struct(in, &data); err != nil {
		return err
	}
	if err = dao.UserGroup.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
			Id: in.GroupId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
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

	if ids, err = s.GetGroupChildrenIDs(ctx, groupId); err != nil {
		return err
	}

	// 删除数据
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
	if t, err = s.GetGroupTree(ctx); err != nil {
		return nil, err
	}
	out = &model.TreeDataOutput{
		TreeData: t.GetTreeData(),
		Total:    uint(t.CountTreeData()),
	}
	return out, nil
}

// 获取分组树
func (s *sUser) GetGroupTree(ctx context.Context) (*tree.Tree, error) {
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
	if t, err = s.GetGroupTree(ctx); err != nil {
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

// 设置分组权限
func (s *sUser) SetupGroupAccess(ctx context.Context, groupId uint, auth_rule_ids []uint) error {
	var (
		err   error
		group *entity.UserGroup
	)

	// 检测分组
	if group, err = s.GetGroup(ctx, groupId); err != nil {
		return err
	}
	if group == nil {
		return gerror.Newf("group is not exists: %d", groupId)
	}
	// 检测权限ID集
	// 待补充...

	var (
		data []*do.UserGroupAccess
	)

	// 组装新增数据
	for _, v := range auth_rule_ids {
		data = append(data, &do.UserGroupAccess{
			GroupId:    groupId,
			AuthRuleId: v,
		})
	}
	if err = dao.UserGroupAccess.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, err = dao.UserGroupAccess.Ctx(ctx).Where(do.UserGroupAccess{
			GroupId: groupId,
		}).Delete(); err != nil {
			return err
		}
		if _, err = dao.UserGroupAccess.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	// 更新授权政策
	service.Oauth().SavePolicy(ctx)

	return nil
}

// 检测分组名称
func (s *sUser) IsGroupNameAvailable(ctx context.Context, name string, notIds ...uint) (bool, error) {
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