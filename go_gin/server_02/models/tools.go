package models

import (
	"time"
)

func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	return str1 + str2
}

// 获取日期
func GetDay() string {
	return time.Now().Format("20060102")
}
