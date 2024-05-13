// 2023/1/3,10:18
package main

import (
	"fmt"
)

//顺序查找
//按顺序比对查找

//在字符串数组中，查找指定字符串
func shunXu(list []string, name string) (index int) {
	index = -1
	for i, v := range list {
		if name == v {
			index = i
		}
	}
	return
}
func main() {
	arr := []string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	fmt.Println(shunXu(arr, "青翼蝠王"))
}
