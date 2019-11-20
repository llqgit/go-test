package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randSeed = int64(1)
	lock     sync.Mutex
)

// 加锁增量种子随机
func IncrSeedRand(max int64) int64 {
	lock.Lock()
	defer lock.Unlock()
	randSeed++
	if randSeed > 100000000 {
		randSeed = 1
	}
	seed := time.Now().UnixNano() + randSeed
	r := rand.New(rand.NewSource(seed))
	return r.Int63n(max)
}
