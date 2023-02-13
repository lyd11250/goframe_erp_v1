package goods

import (
	"context"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sGoods struct {
}

func New() *sGoods {
	return &sGoods{}
}

func init() {
	service.RegisterGoods(New())
}

func (s *sGoods) GetGoodsById(ctx context.Context, in model.GetGoodsByIdInput) (out model.GetGoodsByIdOutput, err error) {
	err = dao.Goods.Ctx(ctx).WherePri(in.GoodsId).Scan(&out)
	return
}

func (s *sGoods) GetGoodsList(ctx context.Context, in model.GetGoodsListInput) (out model.GetGoodsListOutput, err error) {
	err = dao.Goods.Ctx(ctx).WhereLike(dao.Goods.Columns().GoodsName, "%"+*in.GoodsName+"%").Scan(&out.List)
	return
}

func (s *sGoods) AddGoods(ctx context.Context, in model.AddGoodsInput) (out model.AddGoodsOutput, err error) {
	id, err := dao.Goods.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddGoodsOutput{}, err
	}
	out.GoodsId = id
	return
}

func (s *sGoods) UpdateGoods(ctx context.Context, in model.UpdateGoodsInput) (err error) {
	_, err = dao.Goods.Ctx(ctx).OmitNil().Data(in).WherePri(in.GoodsId).Update()
	return
}

func (s *sGoods) GetGoodsUnits(ctx context.Context) (out model.GetGoodsUnitsOutput, err error) {
	column := dao.Goods.Columns().GoodsUnit
	result, err := dao.Goods.Ctx(ctx).Fields(column).Group(column).All()
	if err != nil {
		return model.GetGoodsUnitsOutput{}, err
	}
	for _, v := range result.Array() {
		out.List = append(out.List, v.String())
	}
	return
}
