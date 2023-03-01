package model

import "goframe-erp-v1/internal/model/entity"

type GetGoodsInventoryInput struct {
	GoodsId int64
}

type GetGoodsInventoryOutput struct {
	entity.Inventory
}

type GetInventoryListInput struct {
	Page     int
	PageSize int
}

type GetInventoryListOutput struct {
	Pages int                // 总页数
	Total int                // 总条数
	List  []entity.Inventory // 列表
}

type AddInventoryInput struct {
	entity.Inventory
}

type AddInventoryOutput struct {
	Before entity.Inventory // 入库前库存
	After  entity.Inventory // 入库后库存
}

type ReduceInventoryInput struct {
	GoodsId  int64
	Quantity int
}

type ReduceInventoryOutput struct {
	Before entity.Inventory // 出库前库存
	After  entity.Inventory // 出库后库存
}
