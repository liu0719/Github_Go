// 2022/12/31,20:15
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("now=%v,type:%T\n", now, now)
	//月份会是英文,直接用int转换即是数字月份
	fmt.Print("当前的时间是：", now.Year(), "/", int(now.Month()), "/", now.Day(), " ")
	fmt.Println(now.Hour(), ":", now.Minute(), ":", now.Second())

	//Sprintf会返回一个字符串，而printf直接输出
	//格式化日期
	fmt.Println("------------格式化日期------------")
	fmt.Printf("%d/%d/%d %d:%d:%d\n",
		now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("------------格式化日期2------------")
	//2006 01 02 15 04 05 不可改变  链接字符可以改，数字固定
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//	//time.Sleep(time.Second * 1 / 10)
	//}
	//unix和UnixNano的使用
	fmt.Println(now.Unix())     //unix时间戳
	fmt.Println(now.UnixNano()) //纳秒时间戳
}
