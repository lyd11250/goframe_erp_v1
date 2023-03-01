package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetOrderInfoReq struct {
	g.Meta    `method:"post" path:"/order/info" summary:"获取订单信息" tags:"订单管理"`
	OrderType *int    `json:"orderType" v:"required#订单类型不能为空"`
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
	POrderNo  *string `json:"pOrderNo" v:"required-without:partyId#源订单号不能为空"`
	PartyId   *int64  `json:"partyId" v:"required-without:pOrderId#供应商/客户ID不能为空"`
}

type CreateOrderRes struct {
	Order  map[string]interface{} `json:"order"`
	POrder map[string]interface{} `json:"pOrder"`
}

type CancelCreateOrderReq struct {
	g.Meta    `method:"post" path:"/order/create/cancel" summary:"取消创建订单" tags:"订单管理"`
	OrderType *int    `json:"orderType" v:"required#订单类型不能为空"`
	OrderId   *int64  `json:"orderId" v:"required-without:orderNo#请输入单据ID或单号"`
	OrderNo   *string `json:"orderNo" v:"required-without:orderId#请输入单据ID或单号"`
}

type CancelCreateOrderRes struct {
}
