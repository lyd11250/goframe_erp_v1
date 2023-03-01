package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetGoodsInventoryReq struct {
	g.Meta  `method:"post" path:"/inventory/goods" summary:"获取商品库存" tags:"库存管理"`
	GoodsId int64 `p:"goodsId" v:"required#商品ID不能为空"`
}

type GetGoodsInventoryRes struct {
	entity.Inventory
}

type GetInventoryListReq struct {
	g.Meta   `method:"post" path:"/inventory/list" summary:"获取库存列表" tags:"库存管理"`
	Page     int `p:"page" v:"required#页码不能为空"`
	PageSize int `p:"pageSize" v:"required#每页条数不能为空"`
}

type GetInventoryListRes struct {
	Pages int                `json:"pages"` // 总页数
	Total int                `json:"total"` // 总条数
	List  []entity.Inventory `json:"list"`  // 列表
}

type AddInventoryReq struct {
	g.Meta `method:"post" path:"/inventory/add" summary:"添加库存" tags:"库存管理"`
	entity.Inventory
}

type AddInventoryRes struct {
	Before entity.Inventory `json:"before"` // 入库前库存
	After  entity.Inventory `json:"after"`  // 入库后库存
}

type ReduceInventoryReq struct {
	g.Meta   `method:"post" path:"/inventory/reduce" summary:"减少库存" tags:"库存管理"`
	GoodsId  int64 `p:"goodsId" v:"required#商品ID不能为空"`
	Quantity int   `p:"quantity" v:"required#数量不能为空"`
}

type ReduceInventoryRes struct {
	Before entity.Inventory `json:"before"` // 出库前库存
	After  entity.Inventory `json:"after"`  // 出库后库存
}
