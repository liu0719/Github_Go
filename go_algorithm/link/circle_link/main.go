package main

import "fmt"

//定义结构体结点
type CatNode struct {
	id   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newcat *CatNode) {
	//判断是不是第一只
	if head.next == nil {
		head.id = newcat.id
		head.name = newcat.name
		head.next = head //环装
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newcat
	newcat.next = head
}

//输出环形链表
func ShowCatNode(head *CatNode) {
	fmt.Println("唤醒链表目前:")
	if head.next == nil {
		fmt.Println("空的环形链表")
	}
	temp := head
	for {
		fmt.Print("[", temp.id, temp.name, "]->")
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除唤醒链表中的元素
func DeleCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("空的，，无法删除，，，")
		return head
	}
	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}
	//helper定位到最后
	for {
		if helper.next == head {
			break
		}
		helper = head.next
	}
	//两个以上
	flag := true
	for {
		if temp.next == head { //这里是在比较最后一个
			break
		}
		if temp.id == id {
			if temp == head {
				head = head.next
			}
			//找到，
			helper.next = temp.next
			fmt.Printf("猫猫:%v\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag { //为真则我们上面没有删除
		if temp.id == id {
			helper.next = temp.next
			fmt.Printf("猫猫:%v\n", id)
		} else {
			fmt.Println("没找到该猫")
		}
	}
	return head
}
func main() {
	head := &CatNode{}
	cat := &CatNode{
		id:   1,
		name: "tom",
	}
	InsertCatNode(head, cat)
	cat1 := &CatNode{
		id:   2,
		name: "mary",
	}
	InsertCatNode(head, cat1)
	head = DeleCatNode(head, 4)
	ShowCatNode(head)
}
