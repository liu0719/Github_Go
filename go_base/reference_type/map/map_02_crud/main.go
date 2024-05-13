// 2023/1/4,11:36
package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	//增
	a["no1"] = "松江"
	a["no2"] = "松江"
	a["no3"] = "吴用"
	a["no4"] = "哈哈"
	//删
	//内置函数delete(map,"key)key存在就删除，不存在就不操作，也不会报错
	delete(a, "no1")
	delete(a, "no6")
	//改
	a["no4"] = "哈哈哈"
	fmt.Println(a)
	//如果希望一次性删除所有的key,
	//1遍历逐一删除
	//2直接make一个新的空间
	/*
		a = make(map[string]string)
		fmt.Println(a)
	*/
	//查
	val, ok := a["no2"]
	if ok {
		fmt.Println("有no,值为", val)
	} else {
		fmt.Println("没有")
	}
	fmt.Println("遍历--------------------------")
	//map遍历
	student := make(map[int]map[string]string)
	student[0] = make(map[string]string, 3)
	student[0]["name"] = "tom"
	student[0]["sex"] = "男"
	student[0]["address"] = "北京"
	student[1] = make(map[string]string, 3)
	student[1]["name"] = "mary"
	student[1]["sex"] = "女"
	student[1]["address"] = "上海"
	for i, v1 := range student {
		fmt.Println("i=", i)
		for j, v2 := range v1 {
			fmt.Println(j, v2)
		}
	}
	//map的长度len()
	fmt.Println(len(student))

	//map切片
	fmt.Println("map切片---------------")
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "300"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "500"
	}
	//数组越界了
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "狐狸精"
	//	monsters[2]["age"] = "100"
	//}

	//添加要用到切片的append
	newMonsters := map[string]string{
		"name": "火云邪神",
		"age":  "1200",
	}
	monsters = append(monsters, newMonsters)
	fmt.Println(monsters)
}
