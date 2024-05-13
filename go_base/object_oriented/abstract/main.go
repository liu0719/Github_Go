package main

import (
	"fmt"
	"time"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (a *Account) Deposite(money float64, pwd string) {
	if pwd == a.Pwd {
		if money < 0 {
			fmt.Println("金额不正确...")
			return
		}
		a.Balance += money

		fmt.Println("存款成功...")
	} else {
		fmt.Println("密码错误...")
		return
	}
	time.Sleep(time.Second)
}
func (a *Account) Withdraw(money float64, pwd string) {
	if pwd == a.Pwd {
		if money < 0 {
			fmt.Println("金额不正确...")
			return
		}
		if a.Balance-money < 0 {
			fmt.Println("余额不足...")
			return
		}
		a.Balance -= money
		fmt.Println("取款成功...")
	} else {
		fmt.Println("密码错误...")
		return
	}
	time.Sleep(time.Second)
}
func (a Account) Inquire() {
	fmt.Println("你的余额为:", a.Balance)
	time.Sleep(time.Second)
}
func main() {
	a := Account{
		AccountNo: "账号1",
		Pwd:       "123",
		Balance:   0.0,
	}

	fmt.Println("请输入密码")
	pwd := ""
	fmt.Scanln(&pwd)
	if pwd == a.Pwd {
		fmt.Println("密码正确正在跳转...")
		time.Sleep(time.Second * 2)
		for {
			fmt.Println("请选择服务:1.存款 2.查询余额 3.取款输入 exit退出")
			need := ""
			fmt.Scanln(&need)
			if need == "exit" {
				fmt.Println("退出成功")
				break
			}
			if need == "1" {
				fmt.Println("请输入你要存的数目")
				num := 0.0
				fmt.Scanln(&num)
				a.Deposite(num, pwd)
			}
			if need == "2" {
				a.Inquire()
			}
			if need == "3" {
				fmt.Println("请输入你要取的数目")
				Wnum := 0.0
				fmt.Scanln(&Wnum)
				a.Withdraw(Wnum, pwd)
			}
			fmt.Println("正在跳转到首页")
			time.Sleep(time.Second * 2)
		}
	} else {
		fmt.Println("密码错误")
	}

}
