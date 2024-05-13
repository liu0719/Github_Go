package main

import (
	"errors"
	"fmt"
	"strconv"
)

//使用数组模拟栈
type Stack struct {
	MaxTop int     //栈的最大存放个数
	Top    int     //表示栈顶
	arr    [20]int //数组模拟栈
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

//判断字符是不是运算符
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

//运算方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有误。。。")
	}
	return res
}

//判断优先级
func (this *Stack) Priority(oper int) (res int) {
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}
func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6*5-2"
	//定义索引扫描exp
	index := 0

	//为了配合运算，需要定义变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		//增加一个逻辑，处理多位数

		ch := exp[index : index+1] //切片遍历exp
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) { //说明是符号
			//俩逻辑
			if operStack.Top == -1 { //空栈就直接入栈
				operStack.Push(temp)
			} else {
				/*如果发现operStack栈顶的运算符的优先级大于等于当前准备入栈的运算符
				的优先级就从符号栈pop并从数栈中pop俩数运算后的结果再重新加入数栈
				符号在加入符号栈
				*/
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//运算后的结果重新进入数栈
					numStack.Push(result)
					//当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { //说明是数

			//处理多位数的思路
			// 先定义一个变量KeepNum string做拼接
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			}

			val, _ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}
		//继续扫描
		//先判断index是否已经扫描到最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//运算后的结果重新进入数栈
		numStack.Push(result)
		//当前的符号压入符号栈
	}

	//没有问题则结果就是numStack最后的数
	res, _ := numStack.Pop()
	fmt.Println(res)
}
