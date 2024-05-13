package main

import "fmt"

func main(){
	var n int=10
	fmt.Println(n)
	fmt.Println(&n)
	var ptr *int=&n
	*ptr=9
	fmt.Println(n)
}