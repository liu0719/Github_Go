package main

import (
	"fmt"
	"time"
)

//usb插槽就是现实中的接口。
//你可以把手机,相机, u盘都插在usb插槽上,而不用担心那个插槽是专门插哪个的,原因是做usb插槽的广家和做各种设备的厂家都遵守了统一的规定包括尺寸，排线等等。
//这样的设计需求在Golang编程中也是会大量存在的,我曾经说过,一个程序就是一个世界,在现实世界存在的情况，在程序中也会出现。我们用程序来模拟一下前面的应用场景。
//通用性增强，降低耦合性

//声明/定义一个接口
//可以定义很多方法，不需要实现
//interface不能包含任何变量
//自定义类型要使用时,再根据具体情况把这些方法写出来（实现）
type Usb interface {
	//声明两个没有实现的方法
	start()
	stop()
}

//Phone
type Phone struct {
}

//phon实现usb接口
func (p Phone) start() {
	fmt.Println("手机开始工作。。。")
	time.Sleep(time.Second * 2)
}
func (p Phone) stop() {
	time.Sleep(time.Second * 2)
	fmt.Println("手机停止工作。。。")
}

// camera
type Camera struct {
}

// camera实现USB接口
func (c Camera) start() {
	fmt.Println("相机开始工作。。。")
	time.Sleep(time.Second * 2)
}
func (c Camera) stop() {
	time.Sleep(time.Second * 2)
	fmt.Println("相机停止工作。。。")
}

//computer
type Computer struct {
}

//如果一个结构体，实现了接口中的所有方法，就说这个结构体实现了这个接口
//相反，如果没有实现这个接口中的所有方法的结构体的实例，被当参数传入时，因为没有实现这个接口，就会报错
//computer用working方法调用usb接口实现使与它连接的接口工作
func (c Computer) working(usb Usb) {
	usb.start()
	usb.stop()
}
func main() {
	com := Computer{}
	p := Phone{}
	com.working(p)
	c := Camera{}
	com.working(c)
}
