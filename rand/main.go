package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"sync"
	//"time"
)

// 并发跑随机数，并打印随机结果

func GetRandomNum(max int64) int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(max))
	return num.Int64()
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	m := make(map[int64]int)
	ch := make(chan int64, 10000)

	wg := sync.WaitGroup{}
	wg.Add(100000)

	for c := 0; c < 100000; c++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				num := GetRandomNum(4)
				ch <- num
			}
		}(&wg)
	}

	go func(wg *sync.WaitGroup, ch chan int64) {
		fmt.Println("start wait")
		wg.Wait()
		fmt.Println("start close")
		close(ch)
	}(&wg, ch)

	for num := range ch {
		m[num]++
	}

	fmt.Println("map", m)
}
