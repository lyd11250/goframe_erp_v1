package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetPurchaseOrderReq struct {
	g.Meta  `method:"post" path:"/order/purchase" summary:"获取采购单" tags:"订单管理"`
	OrderId *int64  `p:"orderId" v:"required-without:OrderNo#请输入订单ID或订单号"`
	OrderNo *string `p:"orderNo" v:"required-without:OrderId#请输入订单ID或订单号"`
}

type GetPurchaseOrderRes struct {
	Order entity.PurchaseOrder `json:"order"`
	Items []entity.OrderItem   `json:"items"`
}

type GetPurchaseOrderListReq struct {
	g.Meta   `method:"post" path:"/order/purchase/list" summary:"获取采购单列表" tags:"订单管理"`
	Page     int `p:"page" v:"required#页码不能为空"`
	PageSize int `p:"pageSize" v:"required#每页条数不能为空"`
}

type GetPurchaseOrderListRes struct {
	Pages int                    `json:"pages"`
	Total int                    `json:"total"`
	List  []entity.PurchaseOrder `json:"list"`
}
type CreatePurchaseOrderReq struct {
	g.Meta     `method:"post" path:"/order/purchase/create" summary:"创建采购订单" tags:"订单管理"`
	SupplierId int64 `p:"supplierId" v:"required#供应商ID不能为空"`
}

type CreatePurchaseOrderRes struct {
	entity.PurchaseOrder
}
