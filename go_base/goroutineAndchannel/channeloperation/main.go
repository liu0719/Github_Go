package main

import (
	"fmt"
)

func writerData(numchan chan int) {
	for i := 1; i <= 20000; i++ {
		numchan <- i
		fmt.Println("writer num :", i)
	}
	close(numchan)
}

func readData(numchan chan int, reschan chan map[int]int, exitchan chan bool) {
	for v := range numchan {
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		map1 := make(map[int]int)
		map1[v] = res
		reschan <- map1
		fmt.Println("存入:", map1)

	}
	exitchan <- true
}

func main() {
	numchan := make(chan int, 20000)
	reschan := make(chan map[int]int, 20000)
	exitchan := make(chan bool, 1)
	go writerData(numchan)
	for i := 0; i < 8; i++ {
		go readData(numchan, reschan, exitchan)
	}
lable:
	for {
		for i := 0; i < 8; i++ {
			if <-exitchan {
				fmt.Println("read完成")
				break lable
			}
		}
	}
	close(reschan)

	// // for {
	// // 	v, ok := <-reschan
	// // 	if !ok {
	// // 		break
	// // 	}
	// // 	fmt.Println(v)
	// // }
	fmt.Println("------------------------------")
	for v := range reschan {
		fmt.Println(v)
	}
}
