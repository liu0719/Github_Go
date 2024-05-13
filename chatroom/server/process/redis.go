package process

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

//初始化连接池
func InitPool(adderss string, maxIdle, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", adderss)
		},
	}
}
