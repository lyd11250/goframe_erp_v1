package order

import "goframe-erp-v1/internal/model/entity"

type GetSaleOrderInput struct {
	OrderId *int64
	OrderNo *string
}

type GetSaleOrderOutput struct {
	Order entity.SaleOrder
	Items []entity.OrderItem
}

type GetSaleOrderListInput struct {
	Page     int
	PageSize int
}

type GetSaleOrderListOutput struct {
	Pages int // 总页数
	Total int // 总条数
	List  []entity.SaleOrder
}

type CreateSaleOrderInput struct {
	CustomerId int64
}

type CreateSaleOrderOutput struct {
	entity.SaleOrder
}
