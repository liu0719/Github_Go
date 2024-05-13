package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PutNum(intchan chan int) {
	for i := 0; i < 800000; i++ {
		intchan <- i
	}
	close(intchan)
	wg.Done()
	fmt.Println("放入完成,intchan已经关闭")
}

func primeNum(intchan chan int, primechan chan int) {
	i := 0
	for {

		num, ok := <-intchan
		if !ok { //管道取空了
			break
		}
		//判断是否为素数
		flag := true //假设是素数
		for i := 2; i < (num/2 + 1); i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
			fmt.Println("存入素数 :", num)
			i++
		}
	}
	wg.Done()
}
func main() {
	intchan := make(chan int, 1000)
	primechan := make(chan int, 200000)
	start := time.Now().Unix()

	wg.Add(1)
	go PutNum(intchan)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go primeNum(intchan, primechan)
	}
	wg.Wait()
	close(primechan)
	end := time.Now().Unix()
	i := 0
	for {
		res, ok := <-primechan
		if !ok {
			break
		}
		i++
		//将结果输出
		fmt.Println("素数:", res)
	}
	fmt.Println("1-8000素数共有", i)
	fmt.Println("主线程已退出")
	fmt.Println(end - start)
}
