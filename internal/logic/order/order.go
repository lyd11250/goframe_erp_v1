package order

import (
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/logic/order/impl"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sOrder struct {
	Imap map[int]model.InterfaceOrder
}

func New() *sOrder {
	return &sOrder{
		Imap: map[int]model.InterfaceOrder{
			consts.OrderTypeXSCK: impl.InventoryOrder,
			consts.OrderTypeCGRK: impl.InventoryOrder,
			consts.OrderTypeCGDD: impl.PurchaseOrder,
			consts.OrderTypeXSDD: impl.SaleOrder,
			consts.OrderTypeCGTH: impl.ReturnOrder,
			consts.OrderTypeXSTH: impl.ReturnOrder,
			consts.OrderTypeTHRK: impl.InventoryReturnOrder,
			consts.OrderTypeTHCK: impl.InventoryReturnOrder,
		},
	}
}

func init() {
	service.RegisterOrder(New())
}

func (s *sOrder) GenerateOrderNo(orderType int, time *gtime.Time) (orderNo string) {
	orderNo += consts.OrderPrefixMap[orderType]
	orderNo += time.Format("YmdHisu")
	return
}

func (s *sOrder) RegisterType(orderType int, i model.InterfaceOrder) {
	s.Imap[orderType] = i
}

func (s *sOrder) Type(orderType int) (i model.InterfaceOrder) {
	i = s.Imap[orderType]
	if i == nil {
		panic("订单类型错误")
	}
	return i
}

func (s *sOrder) Prefix(prefix string) (i model.InterfaceOrder) {
	for k, v := range consts.OrderPrefixMap {
		if v == prefix {
			i = s.Imap[k]
			break
		}
	}
	return
}
