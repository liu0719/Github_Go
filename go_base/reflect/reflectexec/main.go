package main

import (
	"fmt"
	"reflect"
)

func reflectfloat(a interface{}) {
	rVal := reflect.ValueOf(a)
	fmt.Println(rVal.Type())
	fmt.Println(rVal.Kind())
	iVal := rVal.Interface()
	fmt.Printf("num =%T\n", iVal)
	num := iVal.(float64)
	fmt.Printf("num =%T\n", num)
}
func main() {
	num := 1.8
	reflectfloat(num)
}
