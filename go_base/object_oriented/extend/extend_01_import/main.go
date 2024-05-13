package main

import (
	"fmt"
)

type Pupil struct {
	Name  string
	Age   int
	Score float64
}

func (p *Pupil) ShowInfo() {
	fmt.Printf("学生姓名：%v，年龄：%v,成绩：%v\n", p.Name, p.Age, p.Score)
}
func (p *Pupil) SetScore(score float64) {
	p.Score = score
}
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试....")
}

type Gur struct {
	Name  string
	Age   int
	Score float64
}

func (p *Gur) GurShowInfo() {
	fmt.Printf("学生姓名：%v，年龄：%v,成绩：%v\n", p.Name, p.Age, p.Score)
}
func (p *Gur) GurSetScore(score float64) {
	p.Score = score
}
func (p *Gur) Gurtesting() {
	fmt.Println("大学生正在考试....")
}
func main() {
	p1 := &Pupil{
		Name: "tom",
		Age:  7,
	}
	p1.testing()
	p1.SetScore(50)
	p1.ShowInfo()
	fmt.Println(*p1)
	p2 := &Gur{
		Name: "bigtom",
		Age:  20,
	}
	p2.Gurtesting()
	p2.GurSetScore(100)
	p2.GurShowInfo()
	fmt.Println(*p2)
}

// 对上面代码的小结
// 1) Pupil 和 Graduate 两个结构体的字段和方法几乎，
// 但是我们却写了相同的代码.复用性不强
// 2)出现代码冗余，而且代码不利于维护，
// 同时也不利于功能的扩展。
