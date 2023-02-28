package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cInventory struct {
}

var Inventory cInventory

func (c *cInventory) AddInventory(ctx context.Context, req *v1.AddInventoryReq) (res *v1.AddInventoryRes, err error) {
	err = service.Inventory().AddInventory(ctx, model.AddInventoryInput{
		Inventory: req.Inventory,
	})
	return
}

func (c *cInventory) UpdateInventory(ctx context.Context, req *v1.UpdateInventoryReq) (res *v1.UpdateInventoryRes, err error) {
	err = service.Inventory().UpdateInventory(ctx, model.UpdateInventoryInput{
		Inventory: req.Inventory,
	})
	return
}

func (c *cInventory) ReduceInventory(ctx context.Context, req *v1.ReduceInventoryReq) (res *v1.ReduceInventoryRes, err error) {
	err = service.Inventory().ReduceInventory(ctx, model.ReduceInventoryInput{
		GoodsId:  req.GoodsId,
		Quantity: req.Quantity,
	})
	return
}

func (c *cInventory) GetGoodsInventory(ctx context.Context, req *v1.GetGoodsInventoryReq) (res *v1.GetGoodsInventoryRes, err error) {
	output, err := service.Inventory().GetGoodsInventory(ctx, model.GetGoodsInventoryInput{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cInventory) GetInventoryStatistic(ctx context.Context, req *v1.GetInventoryStatisticReq) (res *v1.GetInventoryStatisticRes, err error) {
	output, err := service.Inventory().GetInventoryStatistic(ctx)
	if err != nil {
		return
	}
	err = gconv.Struct(output, &res)
	return
}
