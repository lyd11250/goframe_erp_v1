package supplier

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/model/pojo"
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

func (s *sSupplier) GetSupplierList(ctx context.Context) (out pojo.GetSupplierListOutput, err error) {
	err = dao.Supplier.Ctx(ctx).OrderDesc(dao.Supplier.Columns().SupplierStatus).Scan(&out.List)
	return
}

func (s *sSupplier) UpdateSupplier(ctx context.Context, in pojo.UpdateSupplierInput) (err error) {
	_, err = dao.Supplier.Ctx(ctx).OmitNil().Data(in).WherePri(in.SupplierId).Update()
	return
}

func (s *sSupplier) AddSupplier(ctx context.Context, in pojo.AddSupplierInput) (out pojo.AddSupplierOutput, err error) {
	// 判断供应商名称是否存在
	count, err := dao.Supplier.Ctx(ctx).Count(dao.Supplier.Columns().SupplierName, in.SupplierName)
	if err != nil {
		return pojo.AddSupplierOutput{}, err
	}
	if count > 0 {
		return pojo.AddSupplierOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "供应商已存在")
	}
	id, err := dao.Supplier.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return pojo.AddSupplierOutput{}, err
	}
	out = pojo.AddSupplierOutput{SupplierId: id}
	return
}

func (s *sSupplier) GetSupplierById(ctx context.Context, in pojo.GetSupplierByIdInput) (out pojo.GetSupplierByIdOutput, err error) {
	result, err := dao.Supplier.Ctx(ctx).WherePri(in.SupplierId).One()
	if err != nil {
		return out, err
	}
	if result.IsEmpty() {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "供应商不存在")
	}
	err = result.Struct(&out)
	return
}

func (s *sSupplier) CheckSupplierEnabled(ctx context.Context, supplierId int64) (enabled bool, err error) {
	var supplier entity.Supplier
	err = dao.Supplier.Ctx(ctx).WherePri(supplierId).Scan(&supplier)
	if err != nil {
		return false, err
	}
	enabled = supplier.SupplierStatus == consts.StatusEnabled
	return
}
