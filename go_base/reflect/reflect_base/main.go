package main

import (
	"fmt"
	"reflect"
)

//演示反射
func reflectTest01(a interface{}) {
	//通过反射获取传入的变量， type，king,值
	// 1.获取type
	rtype := reflect.TypeOf(a)
	fmt.Println(rtype)

	// 获取到reflect.Value()
	rVal := reflect.ValueOf(a)
	fmt.Println(rVal)

	rrVal := reflect.TypeOf(rVal)
	fmt.Println(rrVal)

	n1 := 20 + rVal.Int()
	fmt.Println(n1)

	//将rval转回interface
	iv := rVal.Interface()
	//将interface{}通过断言转回需要的类型
	num := iv.(int)
	fmt.Printf("%T", num)
}

type Student struct {
	Name string
	Age  int
}

//对结构体的反射
func reflectTest02(a interface{}) {
	// 1.获取type
	rtype := reflect.TypeOf(a)
	fmt.Println(rtype)

	// 获取到reflect.Value()
	rVal := reflect.ValueOf(a)
	fmt.Println(rVal)

	//获取对应的kind
	//rVal.Kind()
	//rtype.Kind()
	fmt.Printf("kind=%v,kind=%v\n", rVal.Kind(), rtype.Kind())

	//将rval转回interface
	iv := rVal.Interface()
	fmt.Printf("%T,,,%v\n", iv, iv)
	//将interface{}通过断言转回需要的类型
	//使用一个待检测的switch的断言形式
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v", stu.Name)
	}

}
func main() {
	num := 100
	reflectTest01(num)

	fmt.Println("_____________________________________")
	stu := Student{
		Name: "tom",
		Age:  100,
	}
	reflectTest02(stu)
}
