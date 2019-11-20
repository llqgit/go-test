package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Rdb *redis.Client
var Nil = redis.Nil

func init() {
	Rdb = NewRedisClient()
}

func NewRedisClient() *redis.Client {
	addr := "localhost:6379"
	pwd := ""
	db := 0
	pool := 0

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
		PoolSize: pool,
	})
	_, err := client.Ping().Result()
	if err != nil {
		//预警
		fmt.Println("redis connect: ", err)
	}
	return client
}
