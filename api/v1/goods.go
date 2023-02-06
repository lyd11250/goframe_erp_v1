package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetGoodsByIdReq struct {
	g.Meta  `path:"/goods/id" method:"post" summary:"通过ID获取商品"`
	GoodsId int64 `json:"goodsId" dc:"商品ID" v:"required#请输入商品ID"`
}

type GetGoodsByIdRes struct {
	entity.Goods
}

type GetGoodsListReq struct {
	g.Meta    `path:"/goods/list" method:"post" summary:"获取所有商品"`
	GoodsName *string `json:"goodsName" dc:"商品名称" v:"required#请输入搜索条件"`
}

type GetGoodsListRes struct {
	List []entity.Goods `json:"list"`
}

type AddGoodsReq struct {
	g.Meta    `path:"/goods/add" method:"post" summary:"新增商品"`
	GoodsName string `json:"goodsName"  dc:"商品名称" v:"required#请输入商品名称"`
	GoodsDesc string `json:"goodsDesc"  dc:"商品描述" v:"required#请输入商品描述"`
	GoodsUnit string `json:"goodsUnit"  dc:"商品单位" v:"required#请输入商品单位"`
	GoodsImg  string `json:"goodsImg"   dc:"商品主图" v:"required#请输入商品主图"`
}

type AddGoodsRes struct {
	GoodsId int64 `json:"goodsId"`
}

type UpdateGoodsReq struct {
	g.Meta      `path:"/goods/update" method:"post" summary:"修改商品"`
	GoodsId     *int64  `json:"goodsId"     dc:"商品ID"`
	GoodsName   *string `json:"goodsName"   dc:"商品名称"`
	GoodsDesc   *string `json:"goodsDesc"   dc:"商品描述"`
	GoodsUnit   *string `json:"goodsUnit"   dc:"商品单位"`
	GoodsImg    *string `json:"goodsImg"    dc:"商品主图"`
	GoodsStatus *int    `json:"goodsStatus" dc:"商品状态"`
}

type UpdateGoodsRes struct {
}
