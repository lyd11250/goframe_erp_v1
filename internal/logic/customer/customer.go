package customer

import (
	"context"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sCustomer struct {
}

func New() *sCustomer {
	return &sCustomer{}
}

func init() {
	service.RegisterCustomer(New())
}

func (s *sCustomer) GetCustomerList(ctx context.Context) (out model.GetCustomerListOutput, err error) {
	err = dao.Customer.Ctx(ctx).OrderDesc(dao.Customer.Columns().CustomerStatus).Scan(&out.List)
	return
}

func (s *sCustomer) UpdateCustomer(ctx context.Context, in model.UpdateCustomerInput) (err error) {
	_, err = dao.Customer.Ctx(ctx).OmitNil().Data(in).WherePri(in.CustomerId).Update()
	return
}

func (s *sCustomer) AddCustomer(ctx context.Context, in model.AddCustomerInput) (out model.AddCustomerOutput, err error) {
	id, err := dao.Customer.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddCustomerOutput{}, err
	}
	out = model.AddCustomerOutput{CustomerId: id}
	return
}

func (s *sCustomer) GetCustomerById(ctx context.Context, in model.GetCustomerByIdInput) (out model.GetCustomerByIdOutput, err error) {
	err = dao.Customer.Ctx(ctx).WherePri(in.CustomerId).Scan(&out)
	return
}
