package main

import (
	"fmt"
	"time"
)

func main() {
	unix := time.Now().Unix()
	filename := fmt.Sprintf("%v-%v", unix, "aa.jpg")
	fmt.Println(filename)
}
