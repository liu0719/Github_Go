// 2023/1/5,11:38
package main

import "fmt"

type A struct {
	Num int
}

func (a A) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(res)
}
func (a A) jisuan2(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}
func (a A) text() {
	fmt.Println(a.Num)
}
func (person Person) GoodPerson() {
	fmt.Println(person.Name, "is a good person.")
}

type Person struct {
	Name string
}

func (a A) getSum(b int, c int) int {
	return b + c
}

type Circle struct {
	radius float64
}

func (cir *Circle) area() float64 {
	return 3.14 * cir.radius * cir.radius
}

type Float float64

func (f Float) print() {
	fmt.Println("f=", f)
}

//编写一个方法改变f的值
func (f *Float) fix() {
	*f = *f + 10.00
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	s.Age++
	str := fmt.Sprintf("name=[%v],age=[%v]", s.Name, s.Age)
	return str
}
func main() {
	a := A{100}
	a.text()
	p := Person{"刘广硕"}
	p.GoodPerson()
	a.jisuan()
	fmt.Println(a.jisuan2(a.Num))
	fmt.Println(a.getSum(3, 4))
	c := Circle{2}
	fmt.Println(c.area())
	f := Float(3.4)
	f.print()
	f.fix()
	f.print()
	stu := Student{"tom", 20}
	fmt.Println(&stu)
}
