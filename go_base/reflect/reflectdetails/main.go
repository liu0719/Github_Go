package main

import (
	"fmt"
	"reflect"
)

func reflect01(a interface{}) {
	rVal := reflect.ValueOf(a)
	//传入地址的话，rVal是一个指针
	fmt.Println(rVal.Kind())
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(100)

}
func main() {
	num := 1
	reflect01(&num)
	fmt.Println(num)

	//可以这样理解rVal.Elem()
	num1 := 9
	ptr := &num1
	num2 := *ptr
	fmt.Println(num2)
}
