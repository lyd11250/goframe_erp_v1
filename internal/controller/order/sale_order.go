package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/api/v1/order"
	orderPojo "goframe-erp-v1/internal/model/pojo/order"
	"goframe-erp-v1/internal/service"
)

func (c *cOrder) GetSaleOrder(ctx context.Context, req *order.GetSaleOrderReq) (res *order.GetSaleOrderRes, err error) {
	var input orderPojo.GetSaleOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetSaleOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) GetSaleOrderList(ctx context.Context, req *order.GetSaleOrderListReq) (res *order.GetSaleOrderListRes, err error) {
	var input orderPojo.GetSaleOrderListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetSaleOrderList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
func (c *cOrder) CreateSaleOrder(ctx context.Context, req *order.CreateSaleOrderReq) (res *order.CreateSaleOrderRes, err error) {
	var input orderPojo.CreateSaleOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return res, err
	}
	output, err := service.Order().CreateSaleOrder(ctx, input)
	if err != nil {
		return res, err
	}
	err = gconv.Struct(output, &res)
	return
}
