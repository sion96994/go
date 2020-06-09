package roughtime

import (
	"sync/atomic"
	"time"
)

var (
	millsecTime atomic.Value
	removeVal = time.Millisecond * 100
)

func init()  {
	t := time.Now().Truncate(removeVal)
	millsecTime.Store(&t)
	go func() {
		for {
			time.Sleep(removeVal)
			t := time.Now().Truncate(removeVal)
			millsecTime.Store(&t)
		}
	}()
}

// 获取当前时间 range(now-100ms, now]
func FloorTimeNow() time.Time {
	t := millsecTime.Load().(*time.Time)
	return *t
}

func CeilingTimeNow() time.Time {
	t := millsecTime.Load().(*time.Time)
	return (*t).Add(removeVal)
}