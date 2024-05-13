package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(j+1, " * ", i+1, " = ", (i+1)*(j+1), " ")
		}
		fmt.Println()
	}

}
