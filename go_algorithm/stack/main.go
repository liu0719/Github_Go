package main

import (
	"errors"
	"fmt"
)

//使用数组模拟栈
type Stack struct {
	MaxTop int    //栈的最大存放个数
	Top    int    //表示栈顶
	arr    [5]int //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {
	//先判断是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满了...")
		return errors.New("栈满了...")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//判断栈是否为空
	if this.Top == -1 {
		fmt.Println("空栈")
		return 0, errors.New("空栈")
	}
	//先取值  ,在Top--
	val = this.arr[this.Top]
	fmt.Println("出栈值:", this.arr[this.Top])
	this.Top--
	return val, nil
}

//遍历栈
func (this *Stack) List() {
	//判断栈是否为空
	if this.Top == -1 {
		fmt.Println("空栈")
		return
	}

	//不为空，就遍历
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%v]=%v\n", i, this.arr[i])
	}

}
func main() {
	stack := &Stack{
		MaxTop: 5,  //最多五个
		Top:    -1, //当栈顶为-1时，表示栈为空
	}

	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	//显示
	stack.List()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.List()
}
