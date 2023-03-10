package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

func (c *cGoods) GetGoodsByName(ctx context.Context, req *v1.GetGoodsByNameReq) (res *v1.GetGoodsByNameRes, err error) {
	output, err := service.Goods().GetGoodsByName(ctx, model.GetGoodsByNameInput{
		GoodsName: req.GoodsName,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.GetGoodsByNameRes{List: output.List}
	return
}

func (c *cGoods) AddGoods(ctx context.Context, req *v1.AddGoodsReq) (res *v1.AddGoodsRes, err error) {
	if req.GoodsPrice <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品价格必须大于0")
	}
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
	if *req.GoodsPrice <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品价格必须大于0")
	}
	input := model.UpdateGoodsInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Goods().UpdateGoods(ctx, input)
	return
}

func (c *cGoods) GetGoodsUnits(ctx context.Context, req *v1.GetGoodsUnitsReq) (res *v1.GetGoodsUnitsRes, err error) {
	output, err := service.Goods().GetGoodsUnits(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetGoodsUnitsRes{List: output.List}
	return
}

func (c *cGoods) GetGoodsSuppliers(ctx context.Context, req *v1.GetGoodsSuppliersReq) (res *v1.GetGoodsSuppliersRes, err error) {
	output, err := service.Goods().GetGoodsSuppliers(ctx, model.GetGoodsSuppliersInput{GoodsId: req.GoodsId})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cGoods) AddGoodsSupplier(ctx context.Context, req *v1.AddGoodsSupplierReq) (res *v1.AddGoodsSupplierRes, err error) {
	if req.SupplyPrice <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "供货价必须大于0")
	}
	input := model.AddGoodsSupplierInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Goods().AddGoodsSupplier(ctx, input)
	return
}

func (c *cGoods) UpdateGoodsSupplier(ctx context.Context, req *v1.UpdateGoodsSupplierReq) (res *v1.UpdateGoodsSupplierRes, err error) {
	if req.SupplyPrice <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "供货价必须大于0")
	}
	input := model.UpdateGoodsSupplierInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Goods().UpdateGoodsSupplier(ctx, input)
	return
}

func (c *cGoods) DeleteGoodsSupplier(ctx context.Context, req *v1.DeleteGoodsSupplierReq) (res *v1.DeleteGoodsSupplierRes, err error) {
	err = service.Goods().
		DeleteGoodsSupplier(ctx, model.DeleteGoodsSupplierInput{
			GoodsId:    req.GoodsId,
			SupplierId: req.SupplierId,
		})
	return
}

func (c *cGoods) GetGoodsListBySupplier(ctx context.Context, req *v1.GetGoodsListBySupplierReq) (res *v1.GetGoodsListBySupplierRes, err error) {
	output, err := service.Goods().GetGoodsListBySupplier(ctx, model.GetGoodsListBySupplierInput{
		SupplierId: req.SupplierId,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}

func (c *cGoods) GetGoodsList(ctx context.Context, req *v1.GetGoodsListReq) (res *v1.GetGoodsListRes, err error) {
	output, err := service.Goods().GetGoodsList(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetGoodsListRes{List: output.List}
	return
}
