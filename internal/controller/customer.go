package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type cCustomer struct {
}

var Customer cCustomer

func (c *cCustomer) GetCustomerList(ctx context.Context, req *v1.GetCustomerListReq) (res *v1.GetCustomerListRes, err error) {
	output, err := service.Customer().GetCustomerList(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCustomerListRes{List: output.List}
	return
}

func (c *cCustomer) UpdateCustomer(ctx context.Context, req *v1.UpdateCustomerReq) (res *v1.UpdateCustomerRes, err error) {
	input := model.UpdateCustomerInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Customer().UpdateCustomer(ctx, input)
	return
}

func (c *cCustomer) AddCustomer(ctx context.Context, req *v1.AddCustomerReq) (res *v1.AddCustomerRes, err error) {
	input := model.AddCustomerInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.Customer().AddCustomer(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.AddCustomerRes{CustomerId: output.CustomerId}
	return
}

func (c *cCustomer) GetCustomerById(ctx context.Context, req *v1.GetCustomerByIdReq) (res *v1.GetCustomerByIdRes, err error) {
	output, err := service.Customer().GetCustomerById(ctx, model.GetCustomerByIdInput{CustomerId: req.CustomerId})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(output, &res)
	return
}
