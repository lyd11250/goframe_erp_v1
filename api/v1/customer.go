package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type CustomerInfo struct {
	CustomerName      string `json:"customerName" dc:"客户名称" v:"required#请输入客户名称"`
	CustomerDesc      string `json:"customerDesc" dc:"客户描述" v:"required#请输入客户描述"`
	CustomerTelephone string `json:"customerTelephone" dc:"客户座机号码" `
	CustomerPhone     string `json:"customerPhone" dc:"客户手机号码" `
	CustomerAddress   string `json:"customerAddress" dc:"客户地址" v:"required#请输入客户地址"`
	CustomerStatus    int8   `json:"customerStatus" dc:"客户状态" v:"required#请输入客户状态"`
}

type GetCustomerListReq struct {
	g.Meta `path:"/customer/list" method:"post" summary:"获取客户列表"`
}

type GetCustomerListRes struct {
	List []entity.Customer `json:"list"`
}

type UpdateCustomerReq struct {
	g.Meta            `path:"/customer/update" method:"post" summary:"修改客户"`
	CustomerId        *int64  `json:"customerId" dc:"客户ID" v:"required#请输入客户ID"`
	CustomerName      *string `json:"customerName" dc:"客户名称"`
	CustomerDesc      *string `json:"customerDesc" dc:"客户描述"`
	CustomerTelephone *string `json:"customerTelephone" dc:"客户座机号码"`
	CustomerPhone     *string `json:"customerPhone" dc:"客户手机号码"`
	CustomerAddress   *string `json:"customerAddress" dc:"客户地址"`
	CustomerStatus    *int8   `json:"customerStatus" dc:"客户状态"`
}

type UpdateCustomerRes struct {
}

type AddCustomerReq struct {
	g.Meta `path:"/customer/add" method:"post" summary:"新增客户"`
	CustomerInfo
}

type AddCustomerRes struct {
	CustomerId int64 `json:"customerId"`
}

type GetCustomerByIdReq struct {
	g.Meta     `path:"/customer/id" method:"post" summary:"通过ID获取客户"`
	CustomerId int64 `json:"customerId"`
}

type GetCustomerByIdRes struct {
	CustomerId int64 `json:"customerId"`
	CustomerInfo
}
