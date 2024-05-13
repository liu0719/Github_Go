// 2023/1/4,10:57
package main

import "fmt"

func main() {
	//map(映射)声明后是不会分布内存的，初始化需要make，分配内存后才能使用
	var a map[string]string
	a = make(map[string]string, 10)
	//key不能一样,重复了会覆盖
	//value可以重复
	//map以前是无序的，现在是有序的了(高版本有序)
	a["no3"] = "松江"
	a["no1"] = "松江"
	a["no2"] = "吴用"
	a["no1"] = "哈哈"
	fmt.Println(a)
	//直接make
	citys := make(map[string]string)
	//直接赋值
	citys1 := map[string]string{
		"haha": "北京",
	}
	fmt.Println(citys)
	fmt.Println(citys1)

	//案例
	student := make(map[int]map[string]string)
	student[0] = make(map[string]string, 3)
	student[0]["name"] = "tom"
	student[0]["sex"] = "男"
	student[0]["address"] = "北京"
	student[1] = make(map[string]string, 3)
	student[1]["name"] = "mary"
	student[1]["sex"] = "女"
	student[1]["address"] = "上海"
	fmt.Println(student[0])
}
