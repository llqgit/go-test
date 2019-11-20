package main

import (
	"fmt"
)

func main() {
	uid := 1234567
	key := "rank:year2019:week42"
	member := fmt.Sprintf("%d", uid)
	value, err := Db.ZRevRank(key, member).Result()
	if err != nil {
		if err == Nil {
			fmt.Println("no data return")
		}
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
