// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-erp-v1/internal/dao/internal"
)

// internalSysAccessDao is internal type for wrapping internal DAO implements.
type internalSysAccessDao = *internal.SysAccessDao

// sysAccessDao is the data access object for table sys_access.
// You can define custom methods on it to extend its functionality as you wish.
type sysAccessDao struct {
	internalSysAccessDao
}

var (
	// SysAccess is globally public accessible object for table sys_access operations.
	SysAccess = sysAccessDao{
		internal.NewSysAccessDao(),
	}
)

// Fill with you ideas below.
