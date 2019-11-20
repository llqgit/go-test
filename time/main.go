package main

import (
	"fmt"
	"github.com/llqgit/go-test/utils"
	"time"
)

func main() {
	//now := time.Now()
	//birthday := time.Date(2001, 9, 1, 0, 0, 0, 0, time.Local)
	//duration := birthday.AddDate(18, 0, 0)
	//fmt.Println(now.UnixNano() > duration.UnixNano())

	fmt.Println(utils.WeeksInYear(time.Now()))

	//fmt.Println(utils.TimeEndOfWeek(time.Now()))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 1)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 2)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 3)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 4)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 5)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 6)))
	//fmt.Println(utils.TimeEndOfWeek(time.Now().AddDate(0, 0, 7)))
}
