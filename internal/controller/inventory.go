package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/service"
)

type cInventory struct {
}

var Inventory cInventory

func (c *cInventory) GetGoodsInventory(ctx context.Context, req *v1.GetGoodsInventoryReq) (res *v1.GetGoodsInventoryRes, err error) {
	output, err := service.Inventory().GetGoodsInventory(ctx, pojo.GetGoodsInventoryInput{GoodsId: req.GoodsId})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cInventory) GetInventoryList(ctx context.Context, req *v1.GetInventoryListReq) (res *v1.GetInventoryListRes, err error) {
	var input pojo.GetInventoryListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Inventory().GetInventoryList(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cInventory) AddInventory(ctx context.Context, req *v1.AddInventoryReq) (res *v1.AddInventoryRes, err error) {
	var input pojo.AddInventoryInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Inventory().AddInventory(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cInventory) ReduceInventory(ctx context.Context, req *v1.ReduceInventoryReq) (res *v1.ReduceInventoryRes, err error) {
	var input pojo.ReduceInventoryInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Inventory().ReduceInventory(ctx, input)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
