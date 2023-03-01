package order

import (
	"goframe-erp-v1/internal/model/entity"
)

type GetInventoryOrderInput struct {
	OrderId *int64
	OrderNo *string
}

type GetInventoryOrderOutput struct {
	Order entity.InventoryOrder
	Items []entity.OrderItem
}

type GetInventoryOrderListInput struct {
	Page      int
	PageSize  int
	OrderType int
}

type GetInventoryOrderListOutput struct {
	Pages int // 总页数
	Total int // 总条数
	List  []entity.InventoryOrder
}

type CreateInventoryOrderInput struct {
	OrderType int
	POrderId  int64
}

type CreateInventoryOrderOutput struct {
	Order  entity.InventoryOrder
	POrder map[string]interface{}
}

type CancelCreateInventoryOrderInput struct {
	OrderId *int64
	OrderNo *string
}
