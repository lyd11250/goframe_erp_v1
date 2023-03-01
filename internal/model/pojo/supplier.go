package pojo

import "goframe-erp-v1/internal/model/entity"

type SupplierInfo struct {
	SupplierName      string // 供应商名称
	SupplierDesc      string // 供应商描述
	SupplierTelephone string // 供应商座机号码
	SupplierPhone     string // 供应商手机号码
	SupplierAddress   string // 供应商地址
	SupplierStatus    string // 供应商状态
}

type GetSupplierListOutput struct {
	List []entity.Supplier
}

type UpdateSupplierInput struct {
	SupplierId        *int64  // 供应商ID
	SupplierName      *string // 供应商名称
	SupplierDesc      *string // 供应商描述
	SupplierTelephone *string // 供应商座机号码
	SupplierPhone     *string // 供应商手机号码
	SupplierAddress   *string // 供应商地址
	SupplierStatus    *string // 供应商状态
}

type AddSupplierInput struct {
	SupplierInfo
}

type AddSupplierOutput struct {
	SupplierId int64
}

type GetSupplierByIdInput struct {
	SupplierId int64
}

type GetSupplierByIdOutput struct {
	entity.Supplier
}
