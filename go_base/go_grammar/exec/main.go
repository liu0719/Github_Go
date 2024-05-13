package main

import "fmt"

// 输入
// scanln
func main() {
	var name string
	var age int
	var sala float32
	var sex string
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sala)
	// fmt.Println("请输入性别")
	// fmt.Scanln(&sex)

	// fmt.Println("我叫",name,"今年",age,"岁了,我的性别是",sex,",我的薪水",sala,"元")

	// 方式二：fmt.Scanf()
	fmt.Println("请输入你的姓名，年龄，薪水，性别，用空格隔开")
	fmt.Scanf("%s %d %f %s", &name, &age, &sala, &sex)
	fmt.Printf("我叫%s，今年%d岁了，我的工资是%f，我的性别是%s", name, age, sala, sex)
}
