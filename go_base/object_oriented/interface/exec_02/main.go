package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}
type SliceHero []Hero

func (h SliceHero) Len() int {
	return len(h)
}

//Less这个方法就是决定你是用什么标准进行排序
//1.按hero的年龄由小到大进行排序
func (h SliceHero) Less(i, j int) bool {
	//返回bool值为真会执行Swap，交换两个值
	return h[i].Age < h[j].Age
	// 修改成对Name排序1
	// return h[i].Name < h[j].Nam
}
func (h SliceHero) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

}
func main() {
	slice1 := []int{1, -4, 99, 4, 8, 6, 23}
	//排序
	//1.冒泡
	//
	sort.Ints(slice1)
	fmt.Println(slice1)

	//对结构体切片进行排序
	// 1.冒泡
	// 2.系统提供的方法

	//测试
	var heroes SliceHero
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄——%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println("-----------------排序前-------------------")
	for _, v := range heroes {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	sort.Sort(heroes)
	fmt.Println("------------------排序后------------------")
	for _, v := range heroes {
		fmt.Printf("%v ", v)
	}
	i := 11
	j := 14
	i, j = j, i
	fmt.Println("i=", i, "j=", j)
}
