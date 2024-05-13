package main

import (
	"fmt"
	"time"
)

func fib(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	start := time.Now()
	fmt.Println(fib(45))

	cost := time.Since(start).Seconds()
	fmt.Println(cost)
}
