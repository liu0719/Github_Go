package main

import (
	"fmt"
	"net/http"
	"time"
)

type myHandler struct {
}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自己创建的处理器处理请\n\n\n\n\n\n\n\n\n\n\n\n\n郭晨阳吃屎")
}
func main() {
	myhandler := myHandler{}
	// http.Handle("/myhandler", &myhandler)

	//创建Server结构
	server := http.Server{
		Addr:        "10.50.142.216:8888",
		Handler:     &myhandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
	// http.ListenAndServe(":8888", nil)
}
