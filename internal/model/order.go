package model

import (
	"context"
	"goframe-erp-v1/internal/model/entity"
)

type InterfaceOrder interface {
	GetOrderInfo(ctx context.Context, in GetOrderInfoInput) (out GetOrderInfoOutput, err error)
	GetOrderList(ctx context.Context, in GetOrderListInput) (out GetOrderListOutput, err error)
	CreateOrder(ctx context.Context, in CreateOrderInput) (out CreateOrderOutput, err error)
	CancelCreateOrder(ctx context.Context, in CancelCreateOrderInput) (err error)
}

type GetOrderInfoInput struct {
	OrderId   *int64
	OrderNo   *string
	OrderType *int
}

type GetOrderInfoOutput struct {
	Order map[string]interface{}
	Items []entity.OrderItem
}

type GetOrderListInput struct {
	Page      int
	PageSize  int
	OrderType int
}

type GetOrderListOutput struct {
	Pages int
	Total int
	List  []map[string]interface{}
}

type CreateOrderInput struct {
	OrderType *int
	POrderNo  *string
	PartyId   *int64
}

type CreateOrderOutput struct {
	Order  map[string]interface{}
	POrder map[string]interface{}
}

type CancelCreateOrderInput struct {
	OrderId *int64
	OrderNo *string
}
