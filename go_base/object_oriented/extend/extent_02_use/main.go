package main

/*
	采用结构体嵌套来实现继承
	可以方法重写（就是子类对父类继承来的方法进行覆盖,就近原则
	赋值时：
	p1 := &Pupil{
		Student{
			Name: "tom",
			Age:  7,
		},
	}
	调用可以省略父类
	p1.testing()
	p1.Name="mary"
	先在p1的Pupil的结构体里找，
	没有的话回去他嵌入的匿名结构体中找
	如果也没有就会报错
	提高代码复用性，扩展性和维护性
*/
import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (s *Student) ShowInfo() {
	fmt.Printf("学生姓名：%v，年龄：%v,成绩：%v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

func (s *Student) testing() {
	fmt.Println("学生正在考试....")
}

//小学生
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试....")
}

//大学生
type Gur struct {
	Student
}

func (p *Gur) Gurtesting() {
	fmt.Println("大学生正在考试....")
}

type A struct {
	Name string
}
type B struct {
	Name string
}
type C struct {
	A
	B
}

//嵌套有名结构体叫组合，不叫继承
//D和A组合
type D struct {
	a A
}
type E struct {
	//结构体内部也可以用指针
	*A
	*B
	//匿名字段，也允许
	//1)如果一个结构体有int类型的匿名字段,就不能第二个。
	//2)如果需要有多个int的字段，则必须给int字段指定名字

	int
}

func main() {
	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	p1 := &Pupil{
		Student{
			Name: "tom",
			Age:  7,
		},
	}
	p1.Student.testing()
	p1.testing()
	p1.SetScore(50)
	p1.ShowInfo()
	fmt.Println(*p1)
	p2 := &Gur{}
	p2.Name = "bigtom"
	p2.Age = 20
	p2.Gurtesting()
	p2.SetScore(100)
	p2.ShowInfo()
	fmt.Println(*p2)

	c := C{}
	//c.name="jack"//error两个Name属性属于同一级，此时不能简写，必须指定A或者B
	// eg:
	c.A.Name = "ikun"
	c.B.Name = "black"

	d := D{
		A{
			Name: "d里面的a",
		},
	}
	fmt.Println(d.a.Name)
	//d.name="jack"//报错，D中的A不是匿名结构体要指定嵌套结构体的字符 eg:
	d.a.Name = "jack"

	e := E{
		//传指针效率更高,使用时要取出，不然会输出地址
		&A{
			Name: "hello",
		},
		&B{
			Name: "helloB",
		},
		int(100),
	}
	e.int = 999
	fmt.Println(e) //adderss
	fmt.Println(*e.A, *e.B)
}
