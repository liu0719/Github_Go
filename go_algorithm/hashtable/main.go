package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) Show() {
	fmt.Printf("链表%v 找到该雇员:%v，姓名:%v\n", this.Id%7, this.Id, this.Name)
}

//EmpLink
//我们这里定义EmpLink不代表头，基地一个环节点就存放数据
type EmpLink struct {
	Head *Emp
}

//添加员工的方法,保证编号是从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil //这是一个辅助指针，在cur的前面
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是空的链表，给emp找到对应位置并插入
	//让cur和emp进行比较然后让pre保持在cur的前面

	for {
		if cur != nil {
			//比较ID
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时,将emp加入到链表的最后
	pre.Next = emp
	emp.Next = cur
}

//显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%v为空!\n", no)
		return
	}
	//遍历当前链表，显示数据
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%v,雇员ID%v,名字:%v=>", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//HashTable
//链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//编写添加雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	//先确定加入到哪一个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//显示所有的雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//标写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到的值就是放入哪个链表的下标
}

func (this *HashTable) FindFById(id int) *Emp {
	//确定在那个链表
	linkno := this.HashFun(id)
	emp := this.LinkArr[linkno].FindById(id)
	return emp
}
func main() {
	var key int
	var id int
	var name string
	var hashtable HashTable
	for {
		fmt.Println("雇员系统菜单")
		fmt.Println("1 添加雇员")
		fmt.Println("2 显示雇员")
		fmt.Println("3 查找雇员")
		fmt.Println("4 退出系统")
		fmt.Println("请输入：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入雇员ID")
			fmt.Scanln(&id)
			fmt.Println("雇员的名字")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case 2:
			hashtable.ShowAll()
		case 3:
			fmt.Println("请输入你要查找的ID")
			fmt.Scanln(&id)
			emp := hashtable.FindFById(id)
			if emp == nil {
				fmt.Println("该ID不存在")
			} else {
				emp.Show()
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误请重新输入")
		}
	}
}
