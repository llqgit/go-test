package utils

import (
	"fmt"
	"time"
)

// 计算函数消耗时间
func TimeCost(start time.Time, name ...string) {
	terminal := time.Since(start)
	if len(name) > 0 {
		fmt.Printf("%vtime cost: %v\n", name, terminal)
	} else {
		fmt.Println("time cost:", terminal)
	}
}

// WeeksInYear 计算时间是当年的第几周（周日为每周最后一天）
func WeeksInYear(t time.Time) (weeks int, year int) {
	yearDay := t.YearDay()
	// 获取今年的年数 eg:2019
	year = t.Year()
	// 判断第一天是周几
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	// 今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		// 如果整除时，为周日。此时应该判断周日是每周第一天，还是每周最后一天
		base := 2
		if (yearDay-firstWeekDays)%7 == 0 {
			base = 1
		}
		week = (yearDay-firstWeekDays)/7 + base
	}
	return week, year
}

// 计算时间到周末的毫秒数
func EndOfWeek(t time.Time) time.Duration {
	weekendNano := TimeEndOfWeek(t).UnixNano()
	return time.Duration(weekendNano - t.UnixNano())
}

// 获取今天最后 24 点时间戳（实际为下一天 0 点）
func TimeEndOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()+1, 0, 0, 0, 0, d.Location())
}

// 获取本周日最后 24 点时间戳
func TimeEndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	temp := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, -1, time.Local)
	if weekday == 0 {
		weekday = 7
	}
	weekend := temp.AddDate(0, 0, 7-weekday)
	return weekend
}

// 获取本月最后 24 点时间戳
func TimeEndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.Local)
}
