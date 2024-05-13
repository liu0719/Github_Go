package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello,world!</h1>")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8888", nil)
}
