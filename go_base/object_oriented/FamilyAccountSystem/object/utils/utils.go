package utils

import (
	"fmt"
	"time"
)

type FamilyAccount struct {
	//声明必要的字段
	//保存用户输入的选项
	key int
	//账户余额
	balance float64
	// //每次收支的金额
	money float64
	// //收支说明
	note string
	//定义标识符记录收支登记的次数
	num int
	//当有收支时，对details进行字符串拼接就可
	details string
}

//写一个工厂模式，返回*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     0,
		balance: 10000.00,
		money:   0.00,
		note:    "",
		num:     0,
		details: "收支\t收支前\t收支金额\t收支后\t说明",
	}
}

//记录
func (f *FamilyAccount) record() {
	fmt.Println("\n-----------------当前收支记录----------------")
	if f.num == 0 {
		fmt.Println("当前没有收支记录，来一笔吧！")
	} else {

		fmt.Println(f.details)

	}
	fmt.Println()
	time.Sleep(time.Second)
}

//收入
func (f *FamilyAccount) inCome() {
	fmt.Println("----登记收入----")
	fmt.Println("请输入本次收入金额:")
	fmt.Scanln(&f.money)
	fmt.Println("请输入本次收入说明:")
	fmt.Scanln(&f.note)
	after := f.balance
	f.balance += f.money
	f.details += fmt.Sprintf("\n收入\t%v\t  %v\t\t%v\t%v", after, f.money, f.balance, f.note)
	f.num++
	time.Sleep(time.Second)
}

//支出
func (f *FamilyAccount) pay() {
	fmt.Println("----登记支出----")
	fmt.Println("请输入本次支出金额:")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足...")
		//break
		return
	}
	fmt.Println("请输入本次支出说明:")
	fmt.Scanln(&f.note)
	after := f.balance
	f.balance -= f.money
	f.details += fmt.Sprintf("\n支出\t%v\t  %v\t\t%v\t%v", after, f.money, f.balance, f.note)
	f.num++
	time.Sleep(time.Second)
}

//退出
func (f *FamilyAccount) exit() string {
	exit := ""
	fmt.Println("你确定要退出吗?(y\\n)")
	fmt.Scanln(&exit)
	return exit
}

//给该结构体绑定相应的方法
func (f *FamilyAccount) Main() {
	for {
		fmt.Println("--------------家庭收支记账系统---------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.登记支出")
		fmt.Println("		4.退出系统")
		fmt.Println("请选择(1-4):")

		fmt.Scanln(&f.key)

		switch f.key {
		case 1:
			f.record()
		case 2:
			f.inCome()
		case 3:
			f.pay()
		case 4:
			ex := f.exit()
			if ex == "y" {
				fmt.Println("退出中...")
				time.Sleep(time.Second)
				fmt.Println("退出成功，欢迎下次使用!...")
				// break lable
				return
			} else if ex == "n" {
				// break
				return
			} else {
				fmt.Println("输入错误")
				// break

			}
		default:
			fmt.Println("输入错误，请重试")
		}
	}
}
