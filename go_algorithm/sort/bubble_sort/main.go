// 2023/1/2,20:18
package main

import (
	"fmt"
)

//排序是将一组数据，按指定顺序进行排序
//排序分类
//1.内部排序，将数据加载到内部存储器进行排序
//(交换排序法)(选择式排序法)(插入式排序法)
//（1） >交换排序

// MaoPaoBtos 1>>冒泡排序法 eg:冒泡由小到大排序
func MaoPaoBtos(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j] = list[j] + list[j+1]
				list[j+1] = list[j] - list[j+1]
				list[j] = list[j] - list[j+1]
			}
		}
	}
	return list
}

//    2>>快速排序法

//2.外部排序，数据量太大，需要借助外部存储进行排序
//(合并排序法)(直接和并排序法)

func main() {
	fmt.Println(MaoPaoBtos([]int{24, 69, 80, 57, 13}))
}
