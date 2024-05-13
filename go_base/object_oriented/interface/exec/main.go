package main

import "fmt"

type Usb interface {
	say()
}

type Stu struct {
}

func (s *Stu) say() {
	fmt.Println("say()")
}
func main() {
	var stu Stu
	//错误！Stu类型没有实现Usb接口，
	//因为say方法绑定的是*Stu类型，Stu并没有实现say方法
	//如果想通过编译，要加上&，取地址符
	var usb Usb = &stu
	usb.say()
	fmt.Println(usb)
}
