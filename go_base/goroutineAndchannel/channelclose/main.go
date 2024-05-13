package main

import "fmt"

func main() {
	intchan := make(chan int, 3)
	intchan <- 100
	intchan <- 199
	close(intchan)
	//这时不能够再写入数据到channel了
	// intchan <- 300 错误
	num := <-intchan
	fmt.Println(num)
	//管道遍历
	intchan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intchan2 <- (i + 1) * 3
	}
	close(intchan2)
	// 遍历管道时，不能用普通的for循环
	for v := range intchan2 {
		fmt.Println(v)
	}
}
