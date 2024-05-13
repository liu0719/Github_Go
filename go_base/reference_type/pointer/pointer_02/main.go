package main

import "fmt"

func main() {
	var i int = 10
	var ptr *int = &i
	var ptr_1 *int = ptr
	fmt.Printf("ptr string:%T,ptr =%v", ptr, ptr)
	fmt.Println()
	fmt.Printf("ptr_1 typr:%T,ptr_1=%v", ptr_1, ptr_1)
}
