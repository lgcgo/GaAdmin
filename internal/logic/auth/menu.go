package auth

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lgcgo/tree"
)

// 添加菜单
func (s *sAuth) CreateMenu(ctx context.Context, in *model.AuthMenuCreateInput) (*entity.AuthMenu, error) {
	var (
		err       error
		available bool
		ent       *entity.AuthMenu
	)

	// 检测父级
	if in.ParentId > 0 {
		var parent *entity.AuthMenu
		parent, err = s.GetMenu(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
	}
	// 标题防重
	if available, err = s.IsMenuTitleAvailable(ctx, in.Title); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("title is already exists: %s", in.Title)
	}

	// 插入数据
	var (
		data     *do.AuthMenu
		insertId int64
	)

	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	if err = dao.AuthRule.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.AuthMenu.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}

	if ent, err = s.GetMenu(ctx, uint(insertId)); err != nil {
		return nil, err
	}

	return ent, nil
}

// 获取菜单
func (s *sAuth) GetMenu(ctx context.Context, menuId uint) (*entity.AuthMenu, error) {
	var (
		ent *entity.AuthMenu
		err error
	)

	// 扫描数据
	if err = dao.AuthMenu.Ctx(ctx).Where(do.AuthMenu{
		Id: menuId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 修改菜单
func (s *sAuth) UpdateMenu(ctx context.Context, in *model.AuthMenuUpdateInput) (*entity.AuthMenu, error) {
	var (
		data      *do.AuthMenu
		ent       *entity.AuthMenu
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetMenu(ctx, in.MenuId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("menu is not exists: %d", in.MenuId)
	}
	// 检测父级
	if in.ParentId > 0 {
		var (
			parent *entity.AuthMenu
			ids    []uint
		)
		parent, err = s.GetMenu(ctx, in.ParentId)
		if parent == nil {
			return nil, gerror.Newf("parent is not exists: %d", in.ParentId)
		}
		for _, v := range ids {
			if in.ParentId == v {
				return nil, gerror.Newf("parent can not be self or child: %d", in.ParentId)
			}
		}
	}
	// 标题防重
	if available, err = s.IsMenuTitleAvailable(ctx, in.Title, []uint{ent.Id}...); err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf("title is already exists: %s", in.Title)
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 更新实体
	if err = dao.AuthMenu.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthMenu.Ctx(ctx).Where(do.AuthMenu{
			Id: in.MenuId,
		}).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	ent, _ = s.GetMenu(ctx, in.MenuId)

	return ent, nil
}

// 删除菜单(硬删除)
func (s *sAuth) DeleteMenu(ctx context.Context, menuId uint) error {
	var (
		ent *entity.AuthMenu
		err error
		ids []uint
	)

	// 扫描数据
	if ent, err = s.GetMenu(ctx, menuId); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("menu is not exists: %d", menuId)
	}
	if ids, err = s.GetMenuChildrenIds(ctx, menuId); err != nil {
		return err
	}

	// 删除数据
	return dao.AuthMenu.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.AuthMenu.Ctx(ctx).WhereIn("id", ids).Delete()
		return err
	})
}

// 获取所有分组
func (s *sAuth) GetAllMenu(ctx context.Context) ([]*entity.AuthMenu, error) {
	var (
		list []*entity.AuthMenu
		err  error
	)

	if err = dao.AuthMenu.Ctx(ctx).Where(do.AuthMenu{
		Status: "normal",
	}).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 获取菜单树数据
func (s *sAuth) GetMenuTreeData(ctx context.Context) (*model.TreeDataOutput, error) {
	var (
		t   *tree.Tree
		out *model.TreeDataOutput
		err error
	)
	if t, err = s.GetMenuTree(ctx); err != nil {
		return nil, err
	}
	out = &model.TreeDataOutput{
		TreeData: t.GetTreeData(),
		Total:    uint(t.CountTreeData()),
	}
	return out, nil
}

// 获取菜单树
func (s *sAuth) GetMenuTree(ctx context.Context) (*tree.Tree, error) {
	var (
		list []*entity.AuthMenu
		out  *tree.Tree
		err  error
	)

	// 获取全部数据
	if list, err = s.GetAllMenu(ctx); err != nil {
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

// 获取菜单子ID集
func (s *sAuth) GetMenuChildrenIds(ctx context.Context, menuId uint) ([]uint, error) {
	var (
		t    *tree.Tree
		err  error
		keys []string
		ids  []uint
	)

	// 获取树对象
	if t, err = s.GetMenuTree(ctx); err != nil {
		return nil, err
	}
	// 获取子健集
	if keys, err = t.GetSpecChildKeys(gconv.String(menuId)); err != nil {
		return nil, err
	}
	// 格式转换
	for _, v := range keys {
		ids = append(ids, gconv.Uint(v))
	}

	return ids, nil
}

// 检测菜单名称
func (s *sAuth) IsMenuTitleAvailable(ctx context.Context, title string, notIds ...uint) (bool, error) {
	var (
		m     = dao.AuthMenu.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.AuthMenu.Columns().Id, v)
	}
	if count, err = m.Where(do.AuthMenu{
		Title: title,
	}).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}
