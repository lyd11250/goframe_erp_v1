// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAccessDao is the data access object for table sys_access.
type SysAccessDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysAccessColumns // columns contains all the column names of Table for convenient usage.
}

// SysAccessColumns defines and stores column names for table sys_access.
type SysAccessColumns struct {
	AccessId    string // 权限ID
	AccessTitle string // 权限标题
	AccessUri   string // 权限路径
}

// sysAccessColumns holds the columns for table sys_access.
var sysAccessColumns = SysAccessColumns{
	AccessId:    "access_id",
	AccessTitle: "access_title",
	AccessUri:   "access_uri",
}

// NewSysAccessDao creates and returns a new DAO object for table data access.
func NewSysAccessDao() *SysAccessDao {
	return &SysAccessDao{
		group:   "default",
		table:   "sys_access",
		columns: sysAccessColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAccessDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAccessDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAccessDao) Columns() SysAccessColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAccessDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAccessDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAccessDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
