package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/api/v1/order"
	orderPojo "goframe-erp-v1/internal/model/pojo/order"
	"goframe-erp-v1/internal/service"
)

func (c *cOrder) GetInventoryOrder(ctx context.Context, req *order.GetInventoryOrderReq) (res *order.GetInventoryOrderRes, err error) {
	var input orderPojo.GetInventoryOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetInventoryOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) GetInventoryOrderList(ctx context.Context, req *order.GetInventoryOrderListReq) (res *order.GetInventoryOrderListRes, err error) {
	var input orderPojo.GetInventoryOrderListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetInventoryOrderList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) CreateInventoryOrder(ctx context.Context, req *order.CreateInventoryOrderReq) (res *order.CreateInventoryOrderRes, err error) {
	var input orderPojo.CreateInventoryOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().CreateInventoryOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
