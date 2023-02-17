package model

import "goframe-erp-v1/internal/model/entity"

type GetGoodsByIdInput struct {
	GoodsId int64
}

type GetGoodsByIdOutput struct {
	entity.Goods
}

type GetGoodsByNameInput struct {
	GoodsName *string
}

type GetGoodsByNameOutput struct {
	List []entity.Goods
}

type AddGoodsInput struct {
	GoodsName  string  // 商品名称
	GoodsDesc  string  // 商品描述
	GoodsUnit  string  // 商品单位
	GoodsPrice float64 // 商品价格
	GoodsImg   string  // 商品主图
}

type AddGoodsOutput struct {
	GoodsId int64
}

type UpdateGoodsInput struct {
	GoodsId     *int64  // 商品ID
	GoodsName   *string // 商品名称
	GoodsDesc   *string // 商品描述
	GoodsUnit   *string // 商品单位
	GoodsPrice  *string // 商品价格
	GoodsImg    *string // 商品主图
	GoodsStatus *int    // 商品状态
}

type GetGoodsUnitsOutput struct {
	List []string
}

type GetGoodsSuppliersInput struct {
	GoodsId int64
}

type GoodsSupplierRel struct {
	SupplierId  int64
	SupplyPrice float64
}

type GetGoodsSuppliersOutput struct {
	List []GoodsSupplierRel
}

type AddGoodsSupplierInput struct {
	GoodsId     int64
	SupplierId  int64
	SupplyPrice float64
}

type UpdateGoodsSupplierInput struct {
	GoodsId     int64
	SupplierId  int64
	SupplyPrice float64
}

type DeleteGoodsSupplierInput struct {
	GoodsId    int64
	SupplierId int64
}

type CheckGoodsEnabledInput struct {
	GoodsId int64
}

type CheckGoodsEnabledOutput struct {
	Enabled bool
}
