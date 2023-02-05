package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type SupplierInfo struct {
	SupplierName      string `json:"supplierName" dc:"供应商名称" v:"required#请输入供应商名称"`
	SupplierDesc      string `json:"supplierDesc" dc:"供应商描述" v:"required#请输入供应商描述"`
	SupplierTelephone string `json:"supplierTelephone" dc:"供应商座机号码" v:"required#请输入供应商座机号码"`
	SupplierPhone     string `json:"supplierPhone" dc:"供应商手机号码" v:"required#请输入供应商手机号码"`
	SupplierAddress   string `json:"supplierAddress" dc:"供应商地址" v:"required#请输入供应商地址"`
	SupplierStatus    int8   `json:"supplierStatus" dc:"供应商状态" v:"required#请输入供应商状态"`
}

type GetSupplierListReq struct {
	g.Meta `path:"/supplier/list" method:"post" summary:"获取供应商列表"`
}

type GetSupplierListRes struct {
	List []entity.Supplier `json:"list"`
}

type UpdateSupplierReq struct {
	g.Meta            `path:"/supplier/update" method:"post" summary:"修改供应商"`
	SupplierId        *int64  `json:"supplierId" dc:"供应商ID" v:"required#请输入供应商ID"`
	SupplierName      *string `json:"supplierName" dc:"供应商名称"`
	SupplierDesc      *string `json:"supplierDesc" dc:"供应商描述"`
	SupplierTelephone *string `json:"supplierTelephone" dc:"供应商座机号码"`
	SupplierPhone     *string `json:"supplierPhone" dc:"供应商手机号码"`
	SupplierAddress   *string `json:"supplierAddress" dc:"供应商地址"`
	SupplierStatus    *int8   `json:"supplierStatus" dc:"供应商状态"`
}

type UpdateSupplierRes struct {
}

type AddSupplierReq struct {
	g.Meta `path:"/supplier/add" method:"post" summary:"新增供应商"`
	SupplierInfo
}

type AddSupplierRes struct {
	SupplierId int64 `json:"supplierId"`
}

type GetSupplierByIdReq struct {
	g.Meta     `path:"/supplier/id" method:"post" summary:"通过ID获取供应商"`
	SupplierId int64 `json:"supplierId"`
}

type GetSupplierByIdRes struct {
	SupplierId int64 `json:"supplierId"`
	SupplierInfo
}
