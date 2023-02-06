package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cGoods struct {
}

var Goods cGoods

func (c *cGoods) GetGoodsById(ctx context.Context, req *v1.GetGoodsByIdReq) (res *v1.GetGoodsByIdRes, err error) {
	output, err := service.Goods().GetGoodsById(ctx, model.GetGoodsByIdInput{GoodsId: req.GoodsId})
	if err != nil {
		return nil, err
	}
	res = &v1.GetGoodsByIdRes{Goods: output.Goods}
	return
}

func (c *cGoods) GetGoodsList(ctx context.Context, req *v1.GetGoodsListReq) (res *v1.GetGoodsListRes, err error) {
	output, err := service.Goods().GetGoodsList(ctx, model.GetGoodsListInput{
		GoodsName: req.GoodsName,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.GetGoodsListRes{List: output.List}
	return
}

func (c *cGoods) AddGoods(ctx context.Context, req *v1.AddGoodsReq) (res *v1.AddGoodsRes, err error) {
	input := model.AddGoodsInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Goods().AddGoods(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.AddGoodsRes{GoodsId: output.GoodsId}
	return
}

func (c *cGoods) UpdateGoods(ctx context.Context, req *v1.UpdateGoodsReq) (res *v1.UpdateGoodsRes, err error) {
	input := model.UpdateGoodsInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Goods().UpdateGoods(ctx, input)
	return
}
