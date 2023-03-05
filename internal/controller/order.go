package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cOrder struct {
}

var Order cOrder

func (c *cOrder) GetOrderInfo(ctx context.Context, req *v1.GetOrderInfoReq) (res *v1.GetOrderInfoRes, err error) {
	var input model.GetOrderInfoInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	var output model.GetOrderInfoOutput
	if input.OrderNo == nil {
		output, err = service.Order().Type(*req.OrderType).GetOrderInfo(ctx, input)
	} else {
		prefix := *req.OrderNo
		prefix = prefix[:4]
		output, err = service.Order().Prefix(prefix).GetOrderInfo(ctx, input)
	}
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) GetOrderList(ctx context.Context, req *v1.GetOrderListReq) (res *v1.GetOrderListRes, err error) {
	var input model.GetOrderListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().Type(*req.OrderType).GetOrderList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	var input model.CreateOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().Type(*req.OrderType).CreateOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) CancelCreateOrder(ctx context.Context, req *v1.CancelCreateOrderReq) (res *v1.CancelCreateOrderRes, err error) {
	var input model.CancelCreateOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	prefix := *req.OrderNo
	prefix = prefix[:4]
	err = service.Order().Prefix(prefix).CancelCreateOrder(ctx, input)
	return
}

func (c *cOrder) InitOrderItem(ctx context.Context, req *v1.InitOrderItemReq) (res *v1.InitOrderItemRes, err error) {
	var input model.InitOrderItemInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	prefix := req.OrderNo[:4]
	err = service.Order().Prefix(prefix).InitOrderItem(ctx, input)
	return
}

func (c *cOrder) CompleteOrder(ctx context.Context, req *v1.CompleteOrderReq) (res *v1.CompleteOrderRes, err error) {
	var input model.CompleteOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	prefix := req.OrderNo[:4]
	err = service.Order().Prefix(prefix).CompleteOrder(ctx, input)
	return
}

func (c *cOrder) CompleteOrderItem(ctx context.Context, req *v1.CompleteOrderItemReq) (res *v1.CompleteOrderItemRes, err error) {
	var input model.CompleteOrderItemInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	prefix := req.OrderNo[:4]
	err = service.Order().Prefix(prefix).CompleteOrderItem(ctx, input)
	return
}

func (c *cOrder) CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	var input model.CancelOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	prefix := req.OrderNo[:4]
	err = service.Order().Prefix(prefix).CancelOrder(ctx, input)
	return
}
