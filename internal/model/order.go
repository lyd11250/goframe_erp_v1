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
	InitOrderItem(ctx context.Context, in InitOrderItemInput) (err error)
	CompleteOrder(ctx context.Context, in CompleteOrderInput) (err error)
	CompleteOrderItem(ctx context.Context, in CompleteOrderItemInput) (err error)
	CancelOrder(ctx context.Context, in CancelOrderInput) (err error)
}

type GetOrderInfoInput struct {
	OrderId   *int64
	OrderNo   *string
	OrderType *int
}

type GetOrderInfoOutput struct {
	Order map[string]interface{}
	Items []entity.OrderItem
	Info  map[string]interface{}
}

type GetOrderListInput struct {
	Page        int
	PageSize    int
	OrderType   int
	OrderStatus *int
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
	Notes     string
}

type CreateOrderOutput struct {
	Order  map[string]interface{}
	POrder map[string]interface{}
}

type CancelCreateOrderInput struct {
	OrderNo *string
}

type InitOrderItemInput struct {
	OrderNo *string
	Items   []entity.OrderItem
}

type CompleteOrderInput struct {
	OrderNo string
	Notes   string
}

type CompleteOrderItemInput struct {
	OrderNo     string
	OrderItemId int64
	Notes       string
}

type CancelOrderInput struct {
	OrderNo string
	Notes   string
}
