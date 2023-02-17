// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-erp-v1/internal/dao/internal"
)

// internalInventoryDao is internal type for wrapping internal DAO implements.
type internalInventoryDao = *internal.InventoryDao

// inventoryDao is the data access object for table inventory.
// You can define custom methods on it to extend its functionality as you wish.
type inventoryDao struct {
	internalInventoryDao
}

var (
	// Inventory is globally public accessible object for table inventory operations.
	Inventory = inventoryDao{
		internal.NewInventoryDao(),
	}
)

// Fill with you ideas below.
