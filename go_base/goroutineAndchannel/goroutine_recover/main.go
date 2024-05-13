package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayhello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
	wg.Done()
}

func test() {

	//这里的defer+recover机制不会影响其它协程执行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test():", err)
		}
	}()
	var mapint map[int]string
	mapint[1] = "goalng" //err没有make，该test携程会出错
	wg.Done()
}
func main() {

	wg.Add(2)
	go sayhello()
	go test()
	wg.Wait()
}
