package order

import "goframe-erp-v1/internal/model/entity"

type GetPurchaseOrderInput struct {
	OrderId *int64
	OrderNo *string
}

type GetPurchaseOrderOutput struct {
	Order entity.PurchaseOrder
	Items []entity.OrderItem
}

type GetPurchaseOrderListInput struct {
	Page     int
	PageSize int
}

type GetPurchaseOrderListOutput struct {
	Pages int // 总页数
	Total int // 总条数
	List  []entity.PurchaseOrder
}

type CreatePurchaseOrderInput struct {
	SupplierId int64
}

type CreatePurchaseOrderOutput struct {
	entity.PurchaseOrder
}
