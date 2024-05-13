package main

import "fmt"

type HeroNode struct {
	id       int
	name     string
	nickname string
	next     *HeroNode //指向下一个结点
}

//编写一个插入结点
//在链表的最后插入
func Insert(head *HeroNode, newhero *HeroNode) {
	//思路
	// 1.先找到链表的最后一个结点
	// 2.创建一个辅助节点，[帮忙]
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newhero
}

//按照ID大小进行排序
func InsertId(head *HeroNode, newhero *HeroNode) {
	//思路
	// 1.先找到链表的最后一个结点
	// 2.创建一个辅助节点，[帮忙]
	temp := head
	for {
		if temp.next == nil {
			break
		} else if temp.next.id > newhero.id { //相同ID也想存的话要加上=
			//一定要先将插入的节点和后面的节点进行连接
			newhero.next = temp.next
			temp.next = newhero.next
			break
		} else if temp.next.id == newhero.id {
			fmt.Println("该Id已经存在")
			return
		}
		temp = temp.next
	}
	temp.next = newhero
}
func Showlist(head *HeroNode) {
	if head.next == nil {
		fmt.Println("空链表----")
		return
	}
	temp := head
	//遍历链表
	for {
		if temp.next == nil {
			break
		}
		fmt.Printf("[%v %v %v]==>", temp.next.id, temp.next.name, temp.next.nickname)
		temp = temp.next
	}
}

//删除一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	for {
		if temp.next == nil {
			fmt.Println("未找到该Id,无法进行删除")
			return
		} else if temp.next.id == id {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}
func main() {
	//创建头节点
	head := &HeroNode{}
	// 2.创建一个新的heroNode
	hero := &HeroNode{
		id:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	Insert(head, hero)
	hero1 := &HeroNode{
		id:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	Insert(head, hero1)
	hero2 := &HeroNode{
		id:       5,
		name:     "林冲",
		nickname: "豹子头",
	}
	Insert(head, hero2)
	hero3 := &HeroNode{
		id:       3,
		name:     "花荣",
		nickname: "小李广",
	}
	InsertId(head, hero3)
	Showlist(head)
	DelHeroNode(head, 2)
	fmt.Println()
	Showlist(head)
}
