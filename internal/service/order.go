// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-erp-v1/internal/model"
)

type (
	IOrder interface {
		CreateOrder(ctx context.Context, in model.CreateOrderInput) (out model.CreateOrderOutput, err error)
		DeleteOrder(ctx context.Context, in model.DeleteOrderInput) (err error)
		GetOrderList(ctx context.Context, in model.GetOrderListInput) (out model.GetOrderListOutput, err error)
		GetOrderById(ctx context.Context, in model.GetOrderByIdInput) (out model.GetOrderByIdOutput, err error)
		SetOrderItem(ctx context.Context, in model.SetOrderItemInput) (err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
