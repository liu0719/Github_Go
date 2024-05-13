package main

import (
	"fmt"
	_ "math"
)

func main() {
	// a := 2.0
	// b := 4.0
	// c := 2.0
	// m := b*b - 4*a*c
	// if m > 0 {
	// 	x1 := (-b + math.Sqrt(m))/2*a
	// 	x2 := (-b - math.Sqrt(m))/2*a
	// 	fmt.Println("俩解")
	// 	fmt.Printf("x1=%v,x2=%v", x1, x2)
	// } else if m == 0 {
	// 	x1 := (-b + math.Sqrt(m))/2*a
	// 	fmt.Println("一个解")
	// 	fmt.Printf("x=%v",x1)
	// } else {
	// 	fmt.Println("无解")
	// }

	// ifelse案例
	// var height int32
	// var money float64
	// var handsome bool
	// fmt.Println("请输入身高")
	// fmt.Scanln(&height)
	// fmt.Println("请输入财富")
	// fmt.Scanln(&money)
	// fmt.Println("请输入长相")
	// fmt.Scanln(&handsome)
	// var a bool=height>180
	// var b bool=money>1e8
	// var c bool
	// if handsome==true{
	// 	c=true
	// }
	// if(a&&b&&c){
	// 	fmt.Println("我一定要嫁给他")
	// }else if a||b||c{
	// 	fmt.Println("勉勉强强")
	// }else{
	// 	fmt.Println("gun")
	// }

	// ifelse不要太多一般不超过三层

	// var time int32
	// var sex int
	// fmt.Println("请输入用时和性别")
	// fmt.Scanf("%d,%d",time,sex)
	// if(time<8){
	// 	if sex=='1'{
	// 		fmt.Println("恭喜你进入男子组决赛")
	// 	}else{
	// 		fmt.Println("恭喜你进入女子组决赛")
	// 	}
	// }else{
	// 		fmt.Println("抱歉，未能进入")
	// 	}

	// 分支语句
	// var month byte
	// var age byte
	// fmt.Println("请输入月份")
	// fmt.Scanln(&month)
	// fmt.Println("请输入你的年龄")
	// fmt.Scanln(&age)

	// if month<=6&&month>=1{
	// 	if age<18{
	// 		fmt.Println("儿童票半价，30元")
	// 	}else if age>60{
	// 		fmt.Println("老人票，15元")
	// 	}else{
	// 		fmt.Println("成人票，60元")
	// 	}
	// }else{
	// 	if age<18&&age>60{
	// 		fmt.Println("非成人票，20元")
	// 	}else{
	// 		fmt.Println("成人票，40元")
	// 	}
	// }

	// switch不用加break；

	// var key byte
	// fmt.Println("请输入字母")
	// fmt.Scanf("%c",&key)

	// switch key {
	// case 'a':
	// 	fmt.Println("这个数是一")
	// case 'b':
	// 	fmt.Println("这个数是二")
	// case 'c':
	// 	fmt.Println("这个数是三")
	// case 'd':
	// 	fmt.Println("这个数是四")
	// case 'e':
	// 	fmt.Println("这个数是五")
	// case 'f':
	// 	fmt.Println("这个数是六")
	// case 'g':
	// 	fmt.Println("这个数是七")

	// default:
	// 	fmt.Println("暂无此人")
	// }

	// case 后可有可以有多个表达式
	// case 也可以不跟表达式 类似于ifelse
	// 也可以对范围进行判断
	


	// switch 穿透 fallthorugh
	// 会执行case的下一个case 默认一层穿透
	// a:=1
	// switch{
	// case a>=1:
	// 	fmt.Println("范围判断")
	// 	fallthrough
	
	// case a>=2:
	// fmt.Println("二")
	
	// }

	var x interface{}
	var y=true
	x=y
	switch i:=x.(type) {
	case nil:
			fmt.Printf("x的类型:%T",i)
	case int:
		fmt.Printf("x的类型:%T",i)
	case float64:
		fmt.Printf("x的类型:%T",i)
	case float32:
		fmt.Printf("x的类型:%T",i)
	case bool,string:
		fmt.Printf("x的类型:%T",i)

	}
}
