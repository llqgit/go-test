package main

import "math/rand"

// 普通固定种子随机（每次执行切换 seed ）
func SeedRand(num int64) int64 {
	return rand.Int63n(num)
}
