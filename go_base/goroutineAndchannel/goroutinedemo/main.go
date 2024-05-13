package main

import (
	"fmt"
	"time"
)

//在主线程(可以理解成进程)中，开启一个goroutine，该协程每隔1秒输出“hello,world"
//在主线程中也每隔一秒输出"hello,golang"，输出10次后,退出程序
//要求主线程和goroutine同时执行

//编写一个函数，每隔一秒输出hello world
func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("test()hell0 world", i)
		time.Sleep(time.Second)
	}

}

//（1）主线程是一个物理线程，直接作用在cpu上，是重量级的，非常耗CPU资源
//（2）协程是从主线程开启的，是轻量级的线程，是逻辑态的，对资源消耗相对较小
//（3）Golang 的协程机制是重要的特点，可以轻松的开启上万个协程。其它编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就突显Golang在并发上的忧势了

func main() { //主线程  遇到go test()会分出一个协程，
	//只要主线程执行完，不管协程有没有执行完，就会退出
	go test() //开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main()hell0 golang", i)
		time.Sleep(time.Second)
	}
}
