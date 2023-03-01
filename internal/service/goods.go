// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-erp-v1/internal/model"
)

type (
	IGoods interface {
		GetGoodsById(ctx context.Context, in model.GetGoodsByIdInput) (out model.GetGoodsByIdOutput, err error)
		GetGoodsByName(ctx context.Context, in model.GetGoodsByNameInput) (out model.GetGoodsByNameOutput, err error)
		AddGoods(ctx context.Context, in model.AddGoodsInput) (out model.AddGoodsOutput, err error)
		UpdateGoods(ctx context.Context, in model.UpdateGoodsInput) (err error)
		GetGoodsUnits(ctx context.Context) (out model.GetGoodsUnitsOutput, err error)
		GetGoodsSuppliers(ctx context.Context, in model.GetGoodsSuppliersInput) (out model.GetGoodsSuppliersOutput, err error)
		AddGoodsSupplier(ctx context.Context, in model.AddGoodsSupplierInput) (err error)
		UpdateGoodsSupplier(ctx context.Context, in model.UpdateGoodsSupplierInput) (err error)
		DeleteGoodsSupplier(ctx context.Context, in model.DeleteGoodsSupplierInput) (err error)
		CheckGoodsEnabled(ctx context.Context, in model.CheckGoodsEnabledInput) (out model.CheckGoodsEnabledOutput, err error)
		GetGoodsListBySupplier(ctx context.Context, in model.GetGoodsListBySupplierInput) (out model.GetGoodsListBySupplierOutput, err error)
		GetGoodsList(ctx context.Context) (out model.GetGoodsListOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
