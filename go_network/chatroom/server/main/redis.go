package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//连接池使用
var pool *redis.Pool

//初始化连接池
func initPool(adderss string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", adderss)
		},
	}
}
