package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type CreateOrderReq struct {
	g.Meta        `path:"/order/create" method:"post" summary:"创建订单" tags:"订单管理"`
	OrderType     int   `json:"orderType" dc:"订单类型" v:"required#请输入订单类型"`
	OrderSupplier int64 `json:"orderSupplier" dc:"供应商ID" v:"required-if:orderType,0#请输入供应商ID"`
	OrderCustomer int64 `json:"orderCustomer" dc:"客户ID" v:"required-if:orderType,1#请输入客户ID"`
}

type CreateOrderRes struct {
	OrderId  int64  `json:"orderId"`
	OrderNum string `json:"orderNum" dc:"订单编号"`
}

type DeleteOrderReq struct {
	g.Meta  `path:"/order/delete" method:"post" summary:"删除订单" tags:"订单管理"`
	OrderId int64 `json:"orderId" dc:"订单ID"`
}

type DeleteOrderRes struct {
}

type GetOrderListReq struct {
	g.Meta    `path:"/order/list" method:"post" summary:"获取订单列表" tags:"订单管理"`
	OrderType int    `json:"orderType" dc:"订单类型" v:"required#请输入订单类型"`
	OrderNum  string `json:"orderNum" dc:"订单编号"`
	PageSize  int    `json:"pageSize" dc:"每页数量"`
	Page      int    `json:"page" dc:"页码"`
}

type GetOrderListRes struct {
	List  []entity.Order `json:"list"`
	Pages int            `json:"pages" dc:"总页数"`
	Total int            `json:"total" dc:"总数量"`
}

type GetOrderByIdReq struct {
	g.Meta  `path:"/order/id" method:"post" summary:"通过订单ID获取订单详情" tags:"订单管理"`
	OrderId int64 `json:"orderId" dc:"订单ID" v:"required#请输入订单ID"`
}

type GetOrderByIdRes struct {
	Order entity.Order       `json:"order"`
	Items []entity.OrderItem `json:"items"`
}

type SetOrderItemReq struct {
	g.Meta  `path:"/order/item/set" method:"post" summary:"设置订单项" tags:"订单管理"`
	OrderId int64              `json:"orderId" dc:"订单ID" v:"required#请输入订单ID"`
	Items   []entity.OrderItem `json:"items" dc:"订单项" v:"required#请输入订单项"`
}

type SetOrderItemRes struct {
}
