// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthMenuDao is the data access object for table auth_menu.
type AuthMenuDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AuthMenuColumns // columns contains all the column names of Table for convenient usage.
}

// AuthMenuColumns defines and stores column names for table auth_menu.
type AuthMenuColumns struct {
	Id       string // ID
	ParentId string // 父ID
	Title    string // 标题
	Remark   string // 备注
	Status   string // 状态
	Weigh    string // 权重
	CreateAt string // 创建日期
	UpdateAt string // 更新日期
}

//  authMenuColumns holds the columns for table auth_menu.
var authMenuColumns = AuthMenuColumns{
	Id:       "id",
	ParentId: "parent_id",
	Title:    "title",
	Remark:   "remark",
	Status:   "status",
	Weigh:    "weigh",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewAuthMenuDao creates and returns a new DAO object for table data access.
func NewAuthMenuDao() *AuthMenuDao {
	return &AuthMenuDao{
		group:   "default",
		table:   "auth_menu",
		columns: authMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthMenuDao) Columns() AuthMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
