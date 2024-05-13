package main

import (
	"fmt"
)

type Kid struct {
	Id   int
	Next *Kid
}

func AddKid(num int) *Kid {
	first := &Kid{} //空节点
	curBoy := &Kid{}
	if num < 1 {
		fmt.Println("num因为正值")
	}
	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		kid := &Kid{
			Id: i,
		}

		if i == 1 {
			first = kid
			curBoy = kid
			curBoy.Next = first
		} else {
			curBoy.Next = kid
			curBoy = kid
			curBoy.Next = first
		}
	}
	return first
}

//显示单向的环形列表【遍历】
func ShowKid(first *Kid) {

	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}
	//创建一个指针[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("ID=%v ->", curBoy.Id)
		//退出的条件
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

//分析思路
// 1.编写一个函数，playgame(first *Kid strataId int countNum int)
//最后编写一个算法按照要求，在环形链表中留下最后一人
func PlayGame(first *Kid, start int, countNum int) {
	// 需要定义辅助指针
	if first.Next == nil {
		fmt.Println("没有小孩，怎么玩?")
		return
	}
	//定义辅助指针，帮助我们删除小孩
	// startId要<=小孩的总数
	tail := first
	//tail移动到最后一个小孩
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//first移动到start（开始的位置）
	for i := 1; i <= start-1; i++ {
		first = first.Next
		tail = tail.Next //tail紧随其后
	}
	//开始
	fmt.Println()
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("ID:%v 出圈->", first.Id)
		first = first.Next
		tail.Next = first
		//相等时，说明只剩下一个小孩
		if first == tail {
			break
		}
	}
	fmt.Printf("ID:%v 出圈->", first.Id)

}
func main() {
	first := AddKid(5)
	ShowKid(first)
	PlayGame(first, 2, 3)
}
