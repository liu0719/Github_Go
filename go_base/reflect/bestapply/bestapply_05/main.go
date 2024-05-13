package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c *Cal) GetSum(name string) string {
	return fmt.Sprintf("%v完成了减法运算,%v - %v = %v", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func TestCal(c interface{}) {
	// typ := reflect.TypeOf(c)
	val := reflect.ValueOf(c)
	num := val.Elem().NumField()
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		val.Elem().Field(i).SetInt(int64(rand.Intn(100)))
	}
	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	res := val.Method(0).Call(params)
	fmt.Println(res)
}
func main() {
	c := &Cal{}
	TestCal(c)
}
