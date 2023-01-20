// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	LogId    int64       `json:"logId"    ` // 日志ID
	UserId   int64       `json:"userId"   ` // 操作者ID
	FromIp   string      `json:"fromIp"   ` // 操作者IP
	Uri      string      `json:"uri"      ` // 访问uri
	LogTime  *gtime.Time `json:"logTime"  ` // 访问时间
	LogType  string      `json:"logType"  ` // 日志级别
	LogError string      `json:"logError" ` // 日志错误信息
}
