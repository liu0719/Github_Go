package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-ch2:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				ch1 <- true
			}
		}
	}()
	go func() {
		j := 'A'
		for {
			select {
			case <-ch1:
				if j > 'Z' {
					return
				}
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++

				ch2 <- true
			}
		}
	}()
	ch2 <- true
	time.Sleep(time.Second * 2)
}
