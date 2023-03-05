// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderItem is the golang structure for table order_item.
type OrderItem struct {
	OrderItemId      int64       `json:"orderItemId"      ` // 订单项ID
	OrderNo          string      `json:"orderNo"          ` // 单号
	GoodsId          int64       `json:"goodsId"          ` // 商品ID
	GoodsName        string      `json:"goodsName"        ` // 商品名称
	Price            float64     `json:"price"            ` // 单价
	Amount           float64     `json:"amount"           ` // 订单项总价格
	Quantity         int         `json:"quantity"         ` // 数量
	Notes            string      `json:"notes"            ` // 备注
	CompleteTime     *gtime.Time `json:"completeTime"     ` // 完成时间
	CompleteUser     int64       `json:"completeUser"     ` // 完成人
	CompleteUserName string      `json:"completeUserName" ` // 完成人姓名
	Status           int         `json:"status"           ` // 订单项状态
}