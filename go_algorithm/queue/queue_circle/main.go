package main

import (
	"errors"
	"fmt"
)

// 使用结构体管理队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.isFull() {
		return errors.New("队列已满")
	}
	//tail不包含最后一个元素
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize //tail向后移一位

	return
}

//判断是否满
func (this *CircleQueue) isFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//显示队列
func (this *CircleQueue) ShowCircleQueue() {
	fmt.Println("队列当前的情况是:")
	size := this.size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("%v=[%v],", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	//先判断队列是否为空
	if this.isEmpty() {
		return -1, errors.New("队列为空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize

	return
}

//判断是否为空
func (this *CircleQueue) isEmpty() bool {
	return this.tail == this.head
}

//
func (this *CircleQueue) size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
func main() {
	//创建队列
	CircleQueue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	for {
		fmt.Println("1.添加数据到队列")
		fmt.Println("2.从队列中取出数据")
		fmt.Println("3.显示队列现在的内容")
		fmt.Println("4.退出")
		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入你要添加的数据")
			var val int
			fmt.Scanln(&val)
			err := CircleQueue.Push(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入成功")
			}
		case 2:
			val, err := CircleQueue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出成功值为:", val)
			}
		case 3:
			CircleQueue.ShowCircleQueue()
		case 4:
			return
		}
	}
}
