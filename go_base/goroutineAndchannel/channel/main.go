package main

import "fmt"

//管道就是队列，先进先出
//多个goroutine操作管道时是安全的，不用锁
// channel是有类型的，一个string的channel只能存放string类型数据

/*
channel是引用类型，要初始化（make)才能用
管道有类型
*/
func main() {
	var a chan int = make(chan int, 2)
	fmt.Println(a)
	//不同于map,容量不会自动增长
	mychan2 := make(chan int, 3)
	fmt.Println(mychan2)
	num := 985
	// 写入数据不能超过容量
	//尾进头出，先进先出
	//超出或无值可取就报错 deadlock
	mychan2 <- num
	mychan2 <- 10
	mychan2 <- 20
	n := <-mychan2
	mychan2 <- 100
	fmt.Println(len(mychan2))
	fmt.Println(cap(mychan2))
	fmt.Println(n)
}
