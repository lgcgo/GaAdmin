// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrgDepartmentDao is the data access object for table org_department.
type OrgDepartmentDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns OrgDepartmentColumns // columns contains all the column names of Table for convenient usage.
}

// OrgDepartmentColumns defines and stores column names for table org_department.
type OrgDepartmentColumns struct {
	Id       string // ID
	ParentId string // 父ID
	Title    string // 标题
	Status   string // 状态
	Weigh    string // 权重
	CreateAt string // 创建日期
	UpdateAt string // 修改日期
}

//  orgDepartmentColumns holds the columns for table org_department.
var orgDepartmentColumns = OrgDepartmentColumns{
	Id:       "id",
	ParentId: "parent_id",
	Title:    "title",
	Status:   "status",
	Weigh:    "weigh",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewOrgDepartmentDao creates and returns a new DAO object for table data access.
func NewOrgDepartmentDao() *OrgDepartmentDao {
	return &OrgDepartmentDao{
		group:   "default",
		table:   "org_department",
		columns: orgDepartmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrgDepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrgDepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrgDepartmentDao) Columns() OrgDepartmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrgDepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrgDepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrgDepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
