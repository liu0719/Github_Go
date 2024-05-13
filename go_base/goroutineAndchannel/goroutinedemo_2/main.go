package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求:现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中。最后显示出来。要求使用goroutine完成
// 思路
// 1.编写一个函数，来计算各个数的阶乘，并放入到map中.2．我们启动的协程多个，统计的将结果放入到 map中
// 3. map应该做出一个全局的.
var (
	myMap = make(map[int]int)
	//声明一个全局的互斥锁
	//sync 是包：synchronized
	//Mutex  互斥
	lock sync.Mutex
)

//计算n的阶乘
func test(n int) {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//放入map

	//加锁
	lock.Lock()
	// //defer
	defer lock.Unlock()
	myMap[n] = res //fatal error: concurrent map writes致命错误，并发写入map

}
func main() {
	//开启多个协程，来完成
	//多个协程操作同一部分空间，存在资源竞争问题，go build -race main.go
	for i := 1; i < 20; i++ {
		go test(i)
	}
	//休眠时间并不合适
	time.Sleep(time.Second * 3)
	lock.Lock()
	defer lock.Unlock()
	for i, v := range myMap {
		fmt.Printf("myMap[%v]=%v\n", i, v)
	}
	//主线程提前结束，协程未执行map为空时提前输出
}
