package timeformat

import (
	"time"
)

func Date(timestamp int64) string {
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t3 := time.Unix(timestamp, 0).In(Loc)
	return t3.Format("2006-01-02 15:04:05")
}
