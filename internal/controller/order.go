package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cOrder struct {
}

var Order cOrder

func (c *cOrder) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	var input model.CreateOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().CreateOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateOrderRes{
		OrderId:  output.OrderId,
		OrderNum: output.OrderNum,
	}
	return
}

func (c *cOrder) DeleteOrder(ctx context.Context, req *v1.DeleteOrderReq) (res *v1.DeleteOrderRes, err error) {
	err = service.Order().DeleteOrder(ctx, model.DeleteOrderInput{OrderId: req.OrderId})
	return
}

func (c *cOrder) GetOrderList(ctx context.Context, req *v1.GetOrderListReq) (res *v1.GetOrderListRes, err error) {
	var input model.GetOrderListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetOrderList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) GetOrderById(ctx context.Context, req *v1.GetOrderByIdReq) (res *v1.GetOrderByIdRes, err error) {
	output, err := service.Order().GetOrderById(ctx, model.GetOrderByIdInput{OrderId: req.OrderId})
	if err != nil {
		return nil, err
	}
	res = &v1.GetOrderByIdRes{}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) SetOrderItem(ctx context.Context, req *v1.SetOrderItemReq) (res *v1.SetOrderItemRes, err error) {
	var input model.SetOrderItemInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Order().SetOrderItem(ctx, input)
	return
}
