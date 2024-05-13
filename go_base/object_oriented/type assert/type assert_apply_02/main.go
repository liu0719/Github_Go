package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func TypeJudge(item ...interface{}) []string {
	arr := []string{}
	for _, v := range item {
		switch v.(type) {
		case bool:
			arr = append(arr, "bool")
		case float64:
			arr = append(arr, "float64")
		case float32:
			arr = append(arr, "float32")
		case int:
			arr = append(arr, "int")
		case int16:
			arr = append(arr, "int16")
		case int32:
			arr = append(arr, "int32")
		case int64:
			arr = append(arr, "int64")
		case string:
			arr = append(arr, "string")
		case Student:
			arr = append(arr, "Student")
		case *Student:
			arr = append(arr, "*Student")
		}
	}
	return arr
}

func main() {
	n1 := 2.2
	var n2 float32 = 23.89
	var n3 int16 = 23
	var n4 int64 = 45
	var n5 bool = false
	var n6 = "terjdk"
	a := Student{"liuxin"}
	b := &Student{"liu"}
	arr := TypeJudge(n1, n2, n3, n4, n5, n6, a, b)
	fmt.Println(arr)
}
