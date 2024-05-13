package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}
}
func main() {
	//先从pool池子中取出一个
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("set", "name", "张三")
	if err != nil {
		fmt.Println("set err:", err)
	}
	res, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
