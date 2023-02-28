package consts

import (
	"time"
)

const (
	LoginExMinute       = 10
	CookieEx            = LoginExMinute * time.Minute
	RedisEx       int64 = LoginExMinute * 60
)

const (
	StatusEnabled  = 1
	StatusDisabled = 0
)

const DefaultPassword = "123456"

const (
	OrderTypePurchase = 0
	OrderTypeSale     = 1

	OrderStatusNew        = 0 // 初始化
	OrderStatusProcessing = 1 // 处理中
	OrderStatusDone       = 2 // 处理完成
	OrderStatusCancel     = 3 // 取消
)

var OrderPrefixMap = map[int]string{
	OrderTypePurchase: "P",
	OrderTypeSale:     "S",
}
