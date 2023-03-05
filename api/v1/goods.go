package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetGoodsByIdReq struct {
	g.Meta  `path:"/goods/id" method:"post" summary:"通过ID获取商品" tags:"商品管理"`
	GoodsId int64 `json:"goodsId" dc:"商品ID" v:"required#请输入商品ID"`
}

type GetGoodsByIdRes struct {
	entity.Goods
}

type GetGoodsByNameReq struct {
	g.Meta    `path:"/goods/name" method:"post" summary:"获取所有商品" tags:"商品管理"`
	GoodsName *string `json:"goodsName" dc:"商品名称" v:"required#请输入商品名称"`
}

type GetGoodsByNameRes struct {
	List []entity.Goods `json:"list"`
}

type AddGoodsReq struct {
	g.Meta     `path:"/goods/add" method:"post" summary:"新增商品" tags:"商品管理"`
	GoodsName  string  `json:"goodsName"  dc:"商品名称" v:"required#请输入商品名称"`
	GoodsDesc  string  `json:"goodsDesc"  dc:"商品描述" v:"required#请输入商品描述"`
	GoodsUnit  string  `json:"goodsUnit"  dc:"商品单位" v:"required#请输入商品单位"`
	GoodsPrice float64 `json:"goodsPrice" dc:"商品价格" v:"float#商品价格必须为数字"`
	GoodsImg   string  `json:"goodsImg"   dc:"商品主图" v:"required#请输入商品主图"`
}

type AddGoodsRes struct {
	GoodsId int64 `json:"goodsId"`
}

type UpdateGoodsReq struct {
	g.Meta      `path:"/goods/update" method:"post" summary:"修改商品" tags:"商品管理"`
	GoodsId     *int64   `json:"goodsId"     dc:"商品ID"`
	GoodsName   *string  `json:"goodsName"   dc:"商品名称"`
	GoodsDesc   *string  `json:"goodsDesc"   dc:"商品描述"`
	GoodsUnit   *string  `json:"goodsUnit"   dc:"商品单位"`
	GoodsPrice  *float64 `json:"goodsPrice"  dc:"商品价格"`
	GoodsImg    *string  `json:"goodsImg"    dc:"商品主图"`
	GoodsStatus *int     `json:"goodsStatus" dc:"商品状态"`
}

type UpdateGoodsRes struct {
}

type GetGoodsUnitsReq struct {
	g.Meta `path:"/goods/units" method:"post" summary:"获取所有商品单位" tags:"商品管理"`
}

type GetGoodsUnitsRes struct {
	List []string `json:"list"`
}

type GetGoodsSuppliersReq struct {
	g.Meta  `path:"/goods/suppliers" method:"post" summary:"获取商品供应商" tags:"商品管理"`
	GoodsId int64 `json:"goodsId" dc:"商品ID" v:"required#请输入商品ID"`
}

type GoodsSupplierRel struct {
	SupplierId  int64   `json:"supplierId"`
	SupplyPrice float64 `json:"supplyPrice"`
}

type GetGoodsSuppliersRes struct {
	List []GoodsSupplierRel `json:"list"`
}

type AddGoodsSupplierReq struct {
	g.Meta      `path:"/goods/supplier/add" method:"post" summary:"新增商品供应商" tags:"商品管理"`
	GoodsId     int64   `json:"goodsId"    dc:"商品ID" v:"required#请输入商品ID"`
	SupplierId  int64   `json:"supplierId" dc:"供应商ID" v:"required#请输入供应商ID"`
	SupplyPrice float64 `json:"supplyPrice" dc:"供应价格" v:"float#供应价格必须为数字"`
}

type AddGoodsSupplierRes struct {
}

type UpdateGoodsSupplierReq struct {
	g.Meta      `path:"/goods/supplier/update" method:"post" summary:"修改商品供应商" tags:"商品管理"`
	GoodsId     int64   `json:"goodsId"    dc:"商品ID" v:"required#请输入商品ID"`
	SupplierId  int64   `json:"supplierId" dc:"供应商ID" v:"required#请输入供应商ID"`
	SupplyPrice float64 `json:"supplyPrice" dc:"供应价格" v:"float#供应价格必须为数字"`
}

type UpdateGoodsSupplierRes struct {
}

type DeleteGoodsSupplierReq struct {
	g.Meta     `path:"/goods/supplier/delete" method:"post" summary:"删除商品供应商" tags:"商品管理"`
	GoodsId    int64 `json:"goodsId"    dc:"商品ID" v:"required#请输入商品ID"`
	SupplierId int64 `json:"supplierId" dc:"供应商ID" v:"required#请输入供应商ID"`
}

type DeleteGoodsSupplierRes struct {
}

type GetGoodsListBySupplierReq struct {
	g.Meta     `path:"/goods/supplier" method:"post" summary:"获取供应商商品列表" tags:"商品管理"`
	SupplierId int64 `json:"supplierId" dc:"供应商ID" v:"required#请输入供应商ID"`
}

type GoodsSupplierRelation struct {
	GoodsId     int64   `json:"goodsId"`
	GoodsName   string  `json:"goodsName"`
	SupplyPrice float64 `json:"supplyPrice"`
}

type GetGoodsListBySupplierRes struct {
	List []GoodsSupplierRelation `json:"list"`
}

type GetGoodsListReq struct {
	g.Meta `path:"/goods/list" method:"post" summary:"获取所有商品" tags:"商品管理"`
}

type GetGoodsListRes struct {
	List []entity.Goods `json:"list"`
}
