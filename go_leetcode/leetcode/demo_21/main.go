package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {
	if list1 == nil && list2 == nil {
		return
	}
	head = list1
	p := list1
	if p != nil {
		for p != nil {
			p = p.Next
		}
		p.Next = list2
		p = list1
	} else {
		p = list2
	}
	// list1和list2连接起来

	for list1 != nil {

		for p != nil {
			if p.Val > list1.Val {
				p.Val, list1.Val = list1.Val, p.Val
			}
			p = p.Next
		}
		p = head
		list1 = list1.Next
	}

	return
}
func main() {
	l11 := &ListNode{
		Val: 1,
	}
	l12 := &ListNode{
		Val: 2,
	}
	l13 := &ListNode{
		Val: 4,
	}
	l11.Next = l12
	l12.Next = l13
	l21 := &ListNode{
		Val: 1,
	}
	l22 := &ListNode{
		Val: 3,
	}
	l23 := &ListNode{
		Val: 4,
	}
	l21.Next = l22
	l22.Next = l23
	head := mergeTwoLists(l11, l21)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
