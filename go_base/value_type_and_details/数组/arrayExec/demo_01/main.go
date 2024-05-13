// 2023/1/1,20:24
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func zimu() []byte {
	var e []byte
	for i := range e {
		e[i] = byte(65 + i)
	}
	return e
}

// ArrMaxI 求数组中最大值
func ArrMaxI(list []int) (max int, index int) {
	index = 0
	max = list[0]
	for i, v := range list {
		if v > max {
			max = v
			index = i
		}
	}
	return
}

func ArrAverage(list []int) (result int, average float64) {

	for _, v := range list {
		result += v
	}
	average = float64(result) / float64(len(list))
	return
}
func randReserve() (arr [5]int) {

	//使每次生成的随机数不一样
	rand.Seed(time.Now().Unix())
	for i := range arr {
		arr[i] = rand.Intn(100)
		fmt.Printf("第%d个随机数是%d\n", i+1, arr[i])
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[len(arr)-1-i] = arr[len(arr)-1-i] + arr[i]
		arr[i] = arr[len(arr)-i-1] - arr[i]
		arr[len(arr)-1-i] = arr[len(arr)-i-1] - arr[i]
	}
	return arr
}
func main() {
	//fmt.Println(ArrMaxI([]int{999, 20, 40, 99, 100}))
	//result, average := ArrAverage([]int{33, 4, 5, 6, 3, 234, 2, 4})
	//fmt.Println(result, average)
	//fmt.Println(randReserve())

	//不借第三个参数交换两值
	//a, b := 10, 20
	//a = a + b //a=30 b=20
	//b = a - b //a=30 b=10
	//a = a - b //a=20 b=10
	//fmt.Printf("a=%v,b=%v\n", a, b)
	fmt.Println(randReserve())
}
