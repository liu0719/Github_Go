package main

import (
	"fmt"
	"time"
)

// 获取当前时间
func GetNOwTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(location)
	fmt.Println(time.Now())
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().String()[:19], location)
	return time
}
func main() {
	now := GetNOwTime()
	fmt.Println(now)
}
