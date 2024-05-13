package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Request)

}
