package order

import (
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/service"
)

type sOrder struct {
}

func New() *sOrder {
	return &sOrder{}
}

func init() {
	service.RegisterOrder(New())
}

func (s *sOrder) generateOrderNo(orderType int, time *gtime.Time) (orderNo string) {
	orderNo += consts.OrderPrefixMap[orderType]
	orderNo += time.Format("YmdHisu")
	return
}
