package main

// 概率计算测试

import (
	"crypto/rand"
	"fmt"
	"github.com/llqgit/go-test/utils"
	"math/big"
	"time"
)

func testIncrRandSeed(max int64, ch chan int64) {
	//defer utils.TimeCost(time.Now())

	//temp := map[int64]int{}

	//for i := int64(0); i < max; i++ {
	num := IncrSeedRand(max)
	//temp[num]++
	ch <- num
	//}
	//fmt.Println(len(temp))
}

func testRandSeed(max int64, ch chan int64) {
	//defer utils.TimeCost(time.Now())

	//temp := map[int64]int{}

	//for i := int64(0); i < max; i++ {
	//num := SeedRand(max)
	num, _ := rand.Int(rand.Reader, big.NewInt(max))
	//temp[num]++
	ch <- num.Int64()
	//}
	//fmt.Println(len(temp))
}

func main() {
	defer utils.TimeCost(time.Now(), "main")
	max := int64(20)
	ch := make(chan int64, 1000000)

	temp := make(map[int64]int)
	for i := int64(0); i < max; i++ {
		//go testIncrRandSeed(max, ch)
		go testRandSeed(max, ch)
	}

	for i := int64(0); i < max; i++ {
		n := <-ch
		lock.Lock()
		temp[n]++
		lock.Unlock()
	}

	//time.Sleep(time.Second * 10)
	for i := range temp {
		fmt.Println(i, temp[i])
	}

	//go func() {
	//	for {
	//		select {}
	//	}
	//}()
}
