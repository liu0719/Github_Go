package main

import (
	"errors"
	"fmt"
)

//使用结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

//添加值
func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	}
	this.rear++ //rear向后移一位
	this.array[this.rear] = val
	return
}

//显示队列
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("%d=[%d],", i, this.array[i])
	}
	fmt.Println()
}

//取出值
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("队列为空")
	}
	this.front++
	val = this.array[this.front]
	return
}
func main() {
	//创建队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入成功")
			}
		case 2:
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出成功值为:", val)
			}
		case 3:
			queue.ShowQueue()
		case 4:
			return
		}
	}
}
