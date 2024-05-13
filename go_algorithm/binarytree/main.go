package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历，先输出根节点,再输出左子树，最后输出左子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%v,,name=%v\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历，先输出左子树,再输出根节点，最后输出右子树
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%v,,name=%v\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//后序遍历，先输出左子树,再输出右子树，最后输出根节点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%v,,name=%v\n", node.No, node.Name)

	}
}
func main() {
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	reight2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	root.Left = left1
	root.Right = right1
	right1.Right = reight2
	fmt.Println("前序遍历")
	PreOrder(root)
	fmt.Println("中序遍历")
	InfixOrder(root)
	fmt.Println("后序遍历")
	PostOrder(root)
}
