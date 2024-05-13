package main

import "fmt"

func main() {
	//管道可以声明为只读或只写

	//1.默认情况下，可读也可写
	chan1 := make(chan int, 10)
	fmt.Println(chan1)

	//声明为只写
	chan2 := make(chan<- int, 10)
	chan2 <- 100
	chan2 <- 2000
	//num:=<-chan2,,错误的，不可取出

	//声明为只读
	chan3 := make(<-chan int, 10)
	num := <-chan3
	fmt.Println(num)

	//chan3<-100,错误的，不可读取

}
