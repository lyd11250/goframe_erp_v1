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
	output, err := service.Order().Type(*req.OrderType).GetOrderInfo(ctx, input)
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
	err = service.Order().Type(*req.OrderType).CancelCreateOrder(ctx, input)
	return
}
