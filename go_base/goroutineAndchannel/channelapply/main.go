package main

import (
	"fmt"
)

//写数据函数
func writerData(intchan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intchan <- i
		fmt.Println("writerData ", i)
	}
	close(intchan)
}

//读数据函数
func readData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Printf("readDat 读到数据=%v\n", v)
	}
	exitchan <- true
	close(exitchan)
}
func main() {
	//创建两个管道
	intchan := make(chan int, 10)
	exitchan := make(chan bool, 50)
	// 开启协程
	go writerData(intchan)
	// go readData(intchan, exitchan)
	// for循环读取exitchan中的
	for {
		if <-exitchan {
			break
		}
	}
}
