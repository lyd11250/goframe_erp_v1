package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cSupplier struct {
}

var Supplier cSupplier

func (c *cSupplier) GetSupplierList(ctx context.Context, req *v1.GetSupplierListReq) (res *v1.GetSupplierListRes, err error) {
	output, err := service.Supplier().GetSupplierList(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetSupplierListRes{List: output.List}
	return
}

func (c *cSupplier) UpdateSupplier(ctx context.Context, req *v1.UpdateSupplierReq) (res *v1.UpdateSupplierRes, err error) {
	input := model.UpdateSupplierInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Supplier().UpdateSupplier(ctx, input)
	return
}

func (c *cSupplier) AddSupplier(ctx context.Context, req *v1.AddSupplierReq) (res *v1.AddSupplierRes, err error) {
	input := model.AddSupplierInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Supplier().AddSupplier(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.AddSupplierRes{SupplierId: output.SupplierId}
	return
}

func (c *cSupplier) GetSupplierById(ctx context.Context, req *v1.GetSupplierByIdReq) (res *v1.GetSupplierByIdRes, err error) {
	output, err := service.Supplier().GetSupplierById(ctx, model.GetSupplierByIdInput{SupplierId: req.SupplierId})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
