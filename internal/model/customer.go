package model

import "goframe-erp-v1/internal/model/entity"

type CustomerInfo struct {
	CustomerName      string // 客户名称
	CustomerDesc      string // 客户描述
	CustomerTelephone string // 客户座机号码
	CustomerPhone     string // 客户手机号码
	CustomerAddress   string // 客户地址
	CustomerStatus    string // 客户状态
}

type GetCustomerListOutput struct {
	List []entity.Customer
}

type UpdateCustomerInput struct {
	CustomerId        *int64  // 客户ID
	CustomerName      *string // 客户名称
	CustomerDesc      *string // 客户描述
	CustomerTelephone *string // 客户座机号码
	CustomerPhone     *string // 客户手机号码
	CustomerAddress   *string // 客户地址
	CustomerStatus    *string // 客户状态
}

type AddCustomerInput struct {
	CustomerInfo
}

type AddCustomerOutput struct {
	CustomerId int64
}

type GetCustomerByIdInput struct {
	CustomerId int64
}

type GetCustomerByIdOutput struct {
	entity.Customer
}
