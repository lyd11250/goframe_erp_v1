package supplier

import (
	"context"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sSupplier struct {
}

func New() *sSupplier {
	return &sSupplier{}
}

func init() {
	service.RegisterSupplier(New())
}

func (s *sSupplier) GetSupplierList(ctx context.Context) (out model.GetSupplierListOutput, err error) {
	err = dao.Supplier.Ctx(ctx).Scan(&out.List)
	return
}

func (s *sSupplier) UpdateSupplier(ctx context.Context, in model.UpdateSupplierInput) (err error) {
	_, err = dao.Supplier.Ctx(ctx).OmitNil().Data(in).WherePri(in.SupplierId).Update()
	return
}

func (s *sSupplier) AddSupplier(ctx context.Context, in model.AddSupplierInput) (out model.AddSupplierOutput, err error) {
	id, err := dao.Supplier.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddSupplierOutput{}, err
	}
	out = model.AddSupplierOutput{SupplierId: id}
	return
}

func (s *sSupplier) GetSupplierById(ctx context.Context, in model.GetSupplierByIdInput) (out model.GetSupplierByIdOutput, err error) {
	err = dao.Supplier.Ctx(ctx).WherePri(in.SupplierId).Scan(&out)
	return
}
