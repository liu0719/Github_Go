// 2023/1/4,14:15
package main

import "fmt"

type stu struct {
	Name    string
	Age     int
	Address string
}

func modify(map1 map[int]int) {
	map1[10] = 100
}
func main() {
	//map是引用类型
	//在一个地方赋值，所有有关的值都会变
	map1 := make(map[int]int)
	map1[0] = 16
	map1[10] = 1
	map1[7] = 15
	map1[18] = 10
	modify(map1)
	fmt.Println(map1)
	//map会自动扩容，不会出现pnic现象
	student := make(map[string]stu)
	stu1 := stu{"tom", 18, "北京"}
	stu2 := stu{"mary", 20, "上海"}
	student["no1"] = stu1
	student["no2"] = stu2
	fmt.Println(student)
	for i, v := range student {
		fmt.Println(i)
		fmt.Println(v.Name)
		fmt.Println(v.Age)
		fmt.Println(v.Address)
	}
}
