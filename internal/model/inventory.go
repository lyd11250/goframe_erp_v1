package model

import "goframe-erp-v1/internal/model/entity"

type GetGoodsInventoryInput struct {
	GoodsId int64
}

type GetGoodsInventoryOutput struct {
	List   []entity.Inventory
	Sum    int
	Amount float64
}

type AddInventoryInput struct {
	entity.Inventory
}

type UpdateInventoryInput struct {
	entity.Inventory
}

type ReduceInventoryInput struct {
	GoodsId  int64
	Quantity int
}

type CheckInventoryInput struct {
	GoodsId  int64
	Quantity int
}

type CheckInventoryOutput struct {
	Enough bool
}

type DeleteInventoryInput struct {
	GoodsId   int64
	GoodsCost float64
}

type GetInventoryStatisticOutput struct {
	Amount  float64
	Sum     int
	Average float64
}
