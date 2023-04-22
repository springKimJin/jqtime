package jqtime

import (
	"time"

	"github.com/shopspring/decimal"
)

// 整数纳秒增加小数点 - 保留精度
func FractionalUnixNanoTime() string {
	nano_time_int := time.Now().UnixNano()
	nano_time_int_str := decimal.NewFromInt(nano_time_int)
	billion_str := decimal.NewFromInt(1000000000)
	return nano_time_int_str.Div(billion_str).String()
}
