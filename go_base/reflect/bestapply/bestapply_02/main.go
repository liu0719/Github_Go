package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
	Sex   string  `json:"sex"`
}

func (m *Monster) Print() {
	fmt.Println("-----start-----")
	fmt.Println(m)
	fmt.Println("----- end------")
}

func (m *Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, score float64, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	// 获取到a对应的类别
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("传入参数不为结构体，已自动退出")
		return
	}
	//获取该结构体题的字段数
	num := val.Elem().NumField()
	fmt.Println("该结构体有", num, "个字段")
	//修改值
	val.Elem().Field(0).SetString("刘广硕")
	//循环输出字段
	for i := 0; i < num; i++ {

		fmt.Printf("字段:%v,值为:%v,", i, val.Elem().Field(i))
		tagval := typ.Elem().Field(i).Tag.Get("json")
		if tagval != "" {
			fmt.Printf("tag标签为%v,\n", tagval)
		} else {
			fmt.Println()
		}
	}
	numMethod := val.NumMethod()
	fmt.Println("该结构体有", numMethod, "个方法")
	// 获取到第二个方法，并调用
	//方法排序是按照函数名字母(ASCII码)进行的排序
	val.Method(1).Call(nil)
	// 调用结构体的第一个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(72))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}
func main() {
	mon := &Monster{
		Name:  "孙悟空",
		Age:   1000,
		Score: 100.2,
		Sex:   "男",
	}
	TestStruct(mon)
}
