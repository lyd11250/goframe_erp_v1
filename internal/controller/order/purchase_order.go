package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/api/v1/order"
	orderPojo "goframe-erp-v1/internal/model/pojo/order"
	"goframe-erp-v1/internal/service"
)

func (c *cOrder) GetPurchaseOrder(ctx context.Context, req *order.GetPurchaseOrderReq) (res *order.GetPurchaseOrderRes, err error) {
	var input orderPojo.GetPurchaseOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetPurchaseOrder(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cOrder) GetPurchaseOrderList(ctx context.Context, req *order.GetPurchaseOrderListReq) (res *order.GetPurchaseOrderListRes, err error) {
	var input orderPojo.GetPurchaseOrderListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Order().GetPurchaseOrderList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
func (c *cOrder) CreatePurchaseOrder(ctx context.Context, req *order.CreatePurchaseOrderReq) (res *order.CreatePurchaseOrderRes, err error) {
	var input orderPojo.CreatePurchaseOrderInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return res, err
	}
	output, err := service.Order().CreatePurchaseOrder(ctx, input)
	if err != nil {
		return res, err
	}
	err = gconv.Struct(output, &res)
	return
}
