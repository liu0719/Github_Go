// 2023/1/4,13:33
package main

import (
	"fmt"
	"sort"
)

func main() {
	list := make(map[int]int, 10)
	list[10] = 100
	list[1] = 13
	list[17] = 15
	list[8] = 4
	list[3] = 9
	fmt.Println(list)
	key := []int{}
	for i := range list {
		key = append(key, i)
	}
	sort.Ints(key)
	fmt.Println(key)
	for _, v := range key {
		fmt.Println(list[v])
	}
}
