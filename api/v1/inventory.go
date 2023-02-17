package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type AddInventoryReq struct {
	g.Meta `path:"/inventory/add" method:"post" summary:"新增库存" tags:"库存管理"`
	entity.Inventory
}

type AddInventoryRes struct {
}

type UpdateInventoryReq struct {
	g.Meta `path:"/inventory/update" method:"post" summary:"修改库存" tags:"库存管理"`
	entity.Inventory
}

type UpdateInventoryRes struct {
}

type ReduceInventoryReq struct {
	g.Meta   `path:"/inventory/reduce" method:"post" summary:"减少库存" tags:"库存管理"`
	GoodsId  int64 `json:"goodsId" dc:"商品ID" v:"required#请输入商品ID"`
	Quantity int   `json:"quantity" dc:"数量" v:"required#请输入数量"`
}

type ReduceInventoryRes struct {
}

type GetGoodsInventoryReq struct {
	g.Meta  `path:"/inventory/goods" method:"post" summary:"通过商品ID获取库存列表" tags:"库存管理"`
	GoodsId int64 `json:"goodsId" dc:"商品ID" v:"required#请输入商品ID"`
}

type GetGoodsInventoryRes struct {
	List   []entity.Inventory `json:"list"`
	Sum    int                `json:"sum"`
	Amount float64            `json:"amount"`
}
