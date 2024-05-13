// 2023/1/1,13:47
package main

import (
	"fmt"
	"strconv"
	"time"
)

func text() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	start := time.Now().Unix()
	text()
	end := time.Now().Unix()
	fmt.Println(end - start)
}
