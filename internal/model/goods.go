package model

import "goframe-erp-v1/internal/model/entity"

type GetGoodsByIdInput struct {
	GoodsId int64
}

type GetGoodsByIdOutput struct {
	entity.Goods
}

type GetGoodsListInput struct {
	GoodsName *string
}

type GetGoodsListOutput struct {
	List []entity.Goods
}

type AddGoodsInput struct {
	GoodsName string // 商品名称
	GoodsDesc string // 商品描述
	GoodsUnit string // 商品单位
	GoodsImg  string // 商品主图
}

type AddGoodsOutput struct {
	GoodsId int64
}

type UpdateGoodsInput struct {
	GoodsId     *int64  // 商品ID
	GoodsName   *string // 商品名称
	GoodsDesc   *string // 商品描述
	GoodsUnit   *string // 商品单位
	GoodsImg    *string // 商品主图
	GoodsStatus *int    // 商品状态
}
