package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetOrderInfoReq struct {
	g.Meta    `method:"post" path:"/order/info" summary:"获取订单信息" tags:"订单管理"`
	OrderType *int    `json:"orderType" v:"required-with:orderId#订单类型不能为空"`
	OrderId   *int64  `json:"orderId" v:"required-without:orderNo#请输入单据ID或单号"`
	OrderNo   *string `json:"orderNo" v:"required-without:orderId#请输入单据ID或单号"`
}

type GetOrderInfoRes struct {
	Order map[string]interface{} `json:"order"`
	Items []entity.OrderItem     `json:"items"`
}

type GetOrderListReq struct {
	g.Meta    `method:"post" path:"/order/list" summary:"获取订单列表" tags:"订单管理"`
	OrderType *int `json:"orderType" v:"required#订单类型不能为空"`
	Page      *int `json:"page" v:"required#页码不能为空"`
	PageSize  *int `json:"pageSize" v:"required#每页数量不能为空"`
}

type GetOrderListRes struct {
	Pages int                      `json:"pages"`
	Total int                      `json:"total"`
	List  []map[string]interface{} `json:"list"`
}

type CreateOrderReq struct {
	g.Meta    `method:"post" path:"/order/create" summary:"创建订单" tags:"订单管理"`
	OrderType *int    `json:"orderType" v:"required#订单类型不能为空"`
	POrderNo  *string `json:"pOrderNo" v:"required-if:orderType,1,orderType,3#源订单号不能为空"`
	PartyId   *int64  `json:"partyId" v:"required-if:orderType,0,orderType,2#供应商/客户ID不能为空"`
	Notes     string  `json:"notes"`
}

type CreateOrderRes struct {
	Order  map[string]interface{} `json:"order"`
	POrder map[string]interface{} `json:"pOrder"`
}

type CancelCreateOrderReq struct {
	g.Meta  `method:"post" path:"/order/create/cancel" summary:"取消创建订单" tags:"订单管理"`
	OrderNo *string `json:"orderNo" v:"required-without:orderId#请输入单号"`
}

type CancelCreateOrderRes struct {
}

type InitOrderItemReq struct {
	g.Meta  `method:"post" path:"/order/item/init" summary:"初始化订单明细" tags:"订单管理"`
	OrderNo string             `json:"orderNo" v:"required#请输入单号"`
	Items   []entity.OrderItem `json:"items" v:"required#请输入明细"`
}

type InitOrderItemRes struct {
}

type CompleteOrderReq struct {
	g.Meta  `method:"post" path:"/order/complete" summary:"完成订单" tags:"订单管理"`
	OrderNo string `json:"orderNo" v:"required#请输入单号"`
	Notes   string `json:"notes"`
}

type CompleteOrderRes struct {
}

type CompleteOrderItemReq struct {
	g.Meta      `method:"post" path:"/order/item/complete" summary:"完成订单项" tags:"订单管理"`
	OrderNo     string `json:"orderNo" v:"required#请输入单号"`
	OrderItemId int64  `json:"orderItemId" v:"required#请输入订单项ID"`
	Notes       string `json:"notes"`
}

type CompleteOrderItemRes struct {
}

type CancelOrderReq struct {
	g.Meta  `method:"post" path:"/order/cancel" summary:"取消订单" tags:"订单管理"`
	OrderNo string `json:"orderNo" v:"required#请输入单号"`
	Notes   string `json:"notes"`
}

type CancelOrderRes struct {
}
