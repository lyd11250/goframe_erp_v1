package consts

import "time"

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
