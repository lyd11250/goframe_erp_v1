// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SaleOrder is the golang structure for table sale_order.
type SaleOrder struct {
	OrderId       int64       `json:"orderId"       ` // 单据ID
	OrderNo       string      `json:"orderNo"       ` // 单号
	CustomerId    int64       `json:"customerId"    ` // 客户ID
	CustomerName  string      `json:"customerName"  ` // 客户名称
	OrderQuantity int         `json:"orderQuantity" ` // 采购数量
	OrderAmount   float64     `json:"orderAmount"   ` // 采购总金额
	CreateTime    *gtime.Time `json:"createTime"    ` // 制单时间
	CreateUser    int64       `json:"createUser"    ` // 制单人
	CompleteTime  *gtime.Time `json:"completeTime"  ` // 完成时间
	CompleteUser  int64       `json:"completeUser"  ` // 操作员
	Notes         string      `json:"notes"         ` // 备注
	OrderStatus   int         `json:"orderStatus"   ` // 单据状态
}
