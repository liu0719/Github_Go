package main

import (
	"fmt"
	"time"
)

func main() {
	//保存用户输入的选项
	key := 0
	//账户余额
	balance := 10000.0
	// //每次收支的金额
	money := 0.0
	// //收支说明
	note := ""
	//定义标识符记录收支登记的次数
	num := 0
	//当有收支时，对details进行字符串拼接就可
	details := "收支\t收支前\t收支金额\t收支后\t说明"

lable:
	for {
		fmt.Println("--------------家庭收支记账系统---------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.登记支出")
		fmt.Println("		4.退出系统")
		fmt.Println("请选择(1-4):")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("\n-----------------当前收支记录----------------")
			if num == 0 {
				fmt.Println("当前没有收支记录，来一笔吧！")
			} else {

				fmt.Println(details)

			}
			fmt.Println()
			time.Sleep(time.Second)
		case 2:
			fmt.Println("----登记收入----")
			fmt.Println("请输入本次收入金额:")
			fmt.Scanln(&money)
			fmt.Println("请输入本次收入说明:")
			fmt.Scanln(&note)
			after := balance
			balance += money
			details += fmt.Sprintf("\n收入\t%v\t  %v\t\t%v\t%v", after, money, balance, note)
			num++
			time.Sleep(time.Second)
		case 3:
			fmt.Println("----登记支出----")
			fmt.Println("请输入本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足...")
				break
			}
			fmt.Println("请输入本次支出说明:")
			fmt.Scanln(&note)
			after := balance
			balance -= money
			details += fmt.Sprintf("\n支出\t%v\t  %v\t\t%v\t%v", after, money, balance, note)
			num++
			time.Sleep(time.Second)
		case 4:
			exit := ""
			fmt.Println("你确定要退出吗?(y\\n)")
			fmt.Scanln(&exit)
			if exit == "y" {
				fmt.Println("退出中...")
				time.Sleep(time.Second)
				fmt.Println("退出成功，欢迎下次使用!...")
				break lable
			} else if exit == "n" {
				break
			} else {
				fmt.Println("输入错误")
				break
			}

		default:
			fmt.Println("输入错误，请重试")
		}
	}
}
