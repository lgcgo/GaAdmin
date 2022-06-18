// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrgMemberDao is the data access object for table org_member.
type OrgMemberDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OrgMemberColumns // columns contains all the column names of Table for convenient usage.
}

// OrgMemberColumns defines and stores column names for table org_member.
type OrgMemberColumns struct {
	Id           string // ID
	UserId       string // 用户ID
	OrgId        string // 公司ID
	Realname     string // 真实名称
	No           string // 工号
	InitPassword string //
	Status       string // 状态
	CreateAt     string // 创建日期
	UpdateAt     string // 更新日期
}

//  orgMemberColumns holds the columns for table org_member.
var orgMemberColumns = OrgMemberColumns{
	Id:           "id",
	UserId:       "user_id",
	OrgId:        "org_id",
	Realname:     "realname",
	No:           "no",
	InitPassword: "init_password",
	Status:       "status",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
}

// NewOrgMemberDao creates and returns a new DAO object for table data access.
func NewOrgMemberDao() *OrgMemberDao {
	return &OrgMemberDao{
		group:   "default",
		table:   "org_member",
		columns: orgMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrgMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrgMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrgMemberDao) Columns() OrgMemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrgMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrgMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrgMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
