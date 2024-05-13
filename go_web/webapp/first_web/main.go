package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自己创建的多路复用器", r.URL.Path)
}
func main() {

	//创建多路复用器
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	//创建路由
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":80", mux)
}
