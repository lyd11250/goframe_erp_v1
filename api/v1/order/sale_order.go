package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetSaleOrderReq struct {
	g.Meta  `method:"post" path:"/order/sale" summary:"获取销售单" tags:"订单管理"`
	OrderId *int64  `p:"orderId" v:"required-without:OrderNo#请输入订单ID或订单号"`
	OrderNo *string `p:"orderNo" v:"required-without:OrderId#请输入订单ID或订单号"`
}

type GetSaleOrderRes struct {
	Order entity.SaleOrder   `json:"order"`
	Items []entity.OrderItem `json:"items"`
}

type GetSaleOrderListReq struct {
	g.Meta   `method:"post" path:"/order/sale/list" summary:"获取销售单列表" tags:"订单管理"`
	Page     int `p:"page" v:"required#页码不能为空"`
	PageSize int `p:"pageSize" v:"required#每页条数不能为空"`
}

type GetSaleOrderListRes struct {
	Pages int                `json:"pages"`
	Total int                `json:"total"`
	List  []entity.SaleOrder `json:"list"`
}
type CreateSaleOrderReq struct {
	g.Meta     `method:"post" path:"/order/sale/create" summary:"创建销售订单" tags:"订单管理"`
	CustomerId int64 `p:"customerId" v:"required#客户ID不能为空"`
}

type CreateSaleOrderRes struct {
	entity.SaleOrder
}
