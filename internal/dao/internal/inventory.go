// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InventoryDao is the data access object for table inventory.
type InventoryDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns InventoryColumns // columns contains all the column names of Table for convenient usage.
}

// InventoryColumns defines and stores column names for table inventory.
type InventoryColumns struct {
	GoodsId  string // 商品ID
	Quantity string // 库存数量
	Amount   string // 总金额
	Price    string // 单位金额
}

// inventoryColumns holds the columns for table inventory.
var inventoryColumns = InventoryColumns{
	GoodsId:  "goods_id",
	Quantity: "quantity",
	Amount:   "amount",
	Price:    "price",
}

// NewInventoryDao creates and returns a new DAO object for table data access.
func NewInventoryDao() *InventoryDao {
	return &InventoryDao{
		group:   "default",
		table:   "inventory",
		columns: inventoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InventoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InventoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InventoryDao) Columns() InventoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InventoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InventoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InventoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
