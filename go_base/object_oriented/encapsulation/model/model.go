package model

import "fmt"

type person struct {
	Name string
	age  int //小写不可导出
	sal  float64
}

// 写一个工厂模式的函数,相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// Age get set
func (p *person) SetAge(age int) {
	if age < 150 && age > 0 {
		p.age = age
	} else {
		fmt.Println("年龄输入不正确。。。")
	}
}
func (p *person) GetAge() int {
	return p.age
}

// sal get set
func (p *person) SetSal(sal float64) {
	p.sal = sal
}
func (p *person) GetSal() float64 {
	return p.sal
}

// exercise
type account struct {
	accountnum string
	pwd        string
	balance    float64
}

func NewAccount(num string, pwd string, bal float64) *account {
	if len(num) > 10 || len(num) < 6 {
		fmt.Println("账号长度在6-10之间哦。。")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度必须六位哦。。")
		return nil
	}
	if bal <= 20 {
		fmt.Println("余额必须大于20。。")
		return nil
	}
	return &account{
		accountnum: num,
		pwd:        pwd,
		balance:    bal,
	}
}

// accountnum set get
func (a *account) SetNum(num string) {
	if len(num) <= 10 && len(num) >= 6 {
		a.accountnum = num
	} else {
		fmt.Println("账号长度在6-10之间哦。。")
	}
}
func (a *account) GetNum() string {
	return a.accountnum
}

// pwd set get
func (a *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		a.pwd = pwd
	} else {
		fmt.Println("密码长度必须六位哦。。")
	}
}
func (a *account) GetPwd() string {
	return a.pwd
}

// bal get set
func (a *account) SetBalance(bal float64) {
	if bal > 20 {
		a.balance = bal
	} else {
		fmt.Println("余额必须大于20。。")
	}
}
func (a *account) GetBalance() float64 {
	return a.balance
}
