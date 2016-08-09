package common

import (
	"time"
)

var SERVICE_NAME = ""

const (
	SUCCESS             = 200 //完全成功
	SUC_PARTIAL         = 201 //部分成功
	SERVER_UNWORKING    = 550 //服务器处于非Working状态
	NO_CONTENT          = 551 //请求结果不存在
	PARAMETER_ERROR     = 552 //参数错误
	INNER_ERROR         = 553 //内部错误
	REQUEST_TOOFAST     = 554 //操作请求超速
	SERVICE_UNSUPPORTED = 555 //服务未启用

)

type (
	PoolParas struct {
		Host               string
		Port               int
		Servermode         int
		PoolSize           int
		Sockettimeout      time.Duration
		PoolTimeout        time.Duration
		IdleTimeout        time.Duration
		IdleCheckFrequency time.Duration
	}
	Result struct {
		Errcode  int32
		Strext   string
		ResValue interface{}
	}
)
