package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	carry := 0
	p := &ListNode{}
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next

		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum %= 10
		newNode := &ListNode{
			Val: sum,
		}

		if head == nil {
			head = newNode
			p = head
		} else {
			p.Next = newNode
			p = p.Next
		}

	}
	if carry > 0 {
		newNode := ListNode{
			Val: carry,
		}
		p.Next = &newNode
	}
	return
}
func main() {
	l11 := &ListNode{
		Val: 2,
	}
	l12 := &ListNode{
		Val: 4,
	}
	l13 := &ListNode{
		Val: 3,
	}
	l11.Next = l12
	l12.Next = l13
	l21 := &ListNode{
		Val: 5,
	}
	l22 := &ListNode{
		Val: 6,
	}
	l23 := &ListNode{
		Val: 4,
	}
	l21.Next = l22
	l22.Next = l23
	head := addTwoNumbers(l11, l21)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
