package util

import (
	"github.com/sion96994/go/roughtime"
	"time"
)

// GetCurrentTs 获取当前时间戳
func GetCurrentTs() int64{
	return time.Now().Unix()
}
// GetZeroTimeStamp 获取一段时间后的零点时间戳
func GetZeroTimeStamp(years, months, days int) int64 {
	t := roughtime.FloorTimeNow()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(years, months, days).Unix()
}

// GetZeroTime 获取一段时间后的零点时间
func GetZeroTime(years, months, days int) time.Time {
	t := roughtime.FloorTimeNow()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(years, months, days)
}

// GetZeroTimeByToday 获取当日零点的时间戳
func GetZeroTimeByToday() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime.Unix()
}

// GetZeroTimeByLastNDay 获取n日前零点的时间戳
func GetZeroTimeByLastNDay(n int64) int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime.Unix() - 86400*n
}

// GetZoreTimeForDate 获取具体某一天零点的时间
func GetZoreTimeForDate(timeStamp int64) int64 {
	year,month,day := time.Unix(timeStamp,0).Date()
	zeroTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return zeroTime.Unix()
}

// GetRemainDay 获取当前时间到每年固定某一天的剩余天数（倒计时天数）
func GetRemainDay(timeStamp int64) (remainDay int32){
	t := time.Unix(timeStamp,0)
	// 获取当年固定日期的0点时间类型
	t1 := time.Date(time.Now().Year(),t.Month(),t.Day(),0,0,0,0,time.Local)
	// 获取当天0点的时间类型
	t2 := time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day(),0,0,0,0,time.Local)
	// t1>t2:当年的目的日期还未经过
	// t1<t2:当年的目的日期已经过了，计算下一年的时间
	// t1=t2:刚好在在目的日期这一天
	if t1.Unix() > t2.Unix() {
		remainDay = int32(t1.Sub(t2).Hours() / 24)
	}else if t1.Unix() < t2.Unix() {
		// 计算出下一年的纪念日
		t3 := time.Date(time.Now().Year()+1,t.Month(),t.Day(),0,0,0,0,time.Local)
		remainDay = int32(t3.Sub(t2).Hours() / 24)
	}else {
		remainDay = int32(t2.Sub(t1).Hours())
	}
	return
}

// GetZeroTsByDate 根据时间 `20200609` 转换零点时间戳
func GetZeroTsByDate(Date string) int64 {
	t, _ := time.ParseInLocation("20060102", Date, time.Local)
	return t.Unix()
}
