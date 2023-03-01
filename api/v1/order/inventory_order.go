package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetInventoryOrderReq struct {
	g.Meta  `method:"post" path:"/order/inventory" summary:"获取出入库单" tags:"订单管理"`
	OrderId *int64  `p:"orderId" v:"required-without:OrderNo#请输入订单ID或订单号"`
	OrderNo *string `p:"orderNo" v:"required-without:OrderId#请输入订单ID或订单号"`
}

type GetInventoryOrderRes struct {
	Order entity.InventoryOrder `json:"order"`
	Items []entity.OrderItem    `json:"items"`
}

type GetInventoryOrderListReq struct {
	g.Meta    `method:"post" path:"/order/inventory/list" summary:"获取出入库单列表" tags:"订单管理"`
	Page      int `p:"page" v:"required#页码不能为空"`
	PageSize  int `p:"pageSize" v:"required#每页条数不能为空"`
	OrderType int `p:"orderType" v:"in:1,3#单据类型错误"`
}

type GetInventoryOrderListRes struct {
	Pages int                     `json:"pages"`
	Total int                     `json:"total"`
	List  []entity.InventoryOrder `json:"list"`
}

type CreateInventoryOrderReq struct {
	g.Meta    `method:"post" path:"/order/inventory/create" summary:"创建出入库单" tags:"订单管理"`
	OrderType int `p:"orderType" v:"in:1,3#单据类型错误"`
	POrderId  int `p:"pOrderId" v:"required#源单据ID不能为空"`
}

type CreateInventoryOrderRes struct {
	Order  entity.InventoryOrder  `json:"order"`
	POrder map[string]interface{} `json:"pOrder"`
}
