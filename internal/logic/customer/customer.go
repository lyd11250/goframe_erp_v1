package customer

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
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
	// 判断客户名称是否存在
	count, err := dao.Customer.Ctx(ctx).Count(dao.Customer.Columns().CustomerName, in.CustomerName)
	if err != nil {
		return model.AddCustomerOutput{}, err
	}
	if count > 0 {
		return model.AddCustomerOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "客户已存在")
	}
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

func (s *sCustomer) CheckCustomerEnabled(ctx context.Context, customerId int64) (enabled bool, err error) {
	var customer entity.Customer
	err = dao.Customer.Ctx(ctx).WherePri(customerId).Scan(&customer)
	if err != nil {
		return false, err
	}
	enabled = customer.CustomerStatus == consts.StatusEnabled
	return
}
