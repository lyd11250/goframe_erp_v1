package consts

import (
	"time"
)

const (
	LoginExMinute       = 30
	CookieEx            = LoginExMinute * time.Minute
	RedisEx       int64 = LoginExMinute * 60
)

const (
	StatusEnabled  = 1
	StatusDisabled = 0
)

const DefaultPassword = "123456"

const (
	OrderTypeXSDD = iota // 销售订单
	OrderTypeXSCK        // 销售出库单
	OrderTypeCGDD        // 采购订单
	OrderTypeCGRK        // 采购入库单
	OrderTypeCGTH        // 采购退货单
	OrderTypeXSTH        // 销售退货单
	OrderTypeTHRK        // 退货入库单
	OrderTypeTHCK        // 退货出库单
)

var OrderPrefixMap = map[int]string{
	OrderTypeCGRK: "CGRK",
	OrderTypeXSCK: "XSCK",
	OrderTypeXSDD: "XSDD",
	OrderTypeCGDD: "CGDD",
	OrderTypeCGTH: "CGTH",
	OrderTypeXSTH: "XSTH",
	OrderTypeTHRK: "THRK",
	OrderTypeTHCK: "THCK",
}

const (
	OrderStatusInit       = iota // 初始状态
	OrderStatusDone              // 已完成
	OrderStatusCancel            // 已取消
	OrderStatusProcessing        // 处理中
)
