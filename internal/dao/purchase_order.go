// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-erp-v1/internal/dao/internal"
)

// internalPurchaseOrderDao is internal type for wrapping internal DAO implements.
type internalPurchaseOrderDao = *internal.PurchaseOrderDao

// purchaseOrderDao is the data access object for table purchase_order.
// You can define custom methods on it to extend its functionality as you wish.
type purchaseOrderDao struct {
	internalPurchaseOrderDao
}

var (
	// PurchaseOrder is globally public accessible object for table purchase_order operations.
	PurchaseOrder = purchaseOrderDao{
		internal.NewPurchaseOrderDao(),
	}
)

// Fill with you ideas below.
