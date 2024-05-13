package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("aa.xlsx", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}
