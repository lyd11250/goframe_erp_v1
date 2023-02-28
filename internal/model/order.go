package model

import "goframe-erp-v1/internal/model/entity"

type CreateOrderInput struct {
	OrderType     int   // 订单类型
	OrderSupplier int64 // 供应商ID
	OrderCustomer int64 // 客户ID
}

type CreateOrderOutput struct {
	OrderId  int64
	OrderNum string
}

type DeleteOrderInput struct {
	OrderId int64 // 订单ID
}

type GetOrderListInput struct {
	OrderType int    // 订单类型
	OrderNum  string // 订单编号
	PageSize  int    // 每页数量
	Page      int    // 页码
}

type GetOrderListOutput struct {
	List  []entity.Order
	Pages int
	Total int
}

type GetOrderByIdInput struct {
	OrderId int64 // 订单ID
}

type GetOrderByIdOutput struct {
	Order entity.Order
	Items []entity.OrderItem
}

type SetOrderItemInput struct {
	OrderId int64              // 订单ID
	Items   []entity.OrderItem // 订单项
}
