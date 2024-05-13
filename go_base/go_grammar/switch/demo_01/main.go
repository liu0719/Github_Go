package main

import "fmt"

func main() {
	// var a byte
	// fmt.Println("请输入一个字母")
	// fmt.Scanf("%c",&a)
	// switch a {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// default:
	// 	fmt.Println("未知字符")
	// }

	// var num int
	// fmt.Scanf("%d",&num)

	// switch  {
	// case  num>=60&&num<=100:
	// 	fmt.Println("及格")
	// case num>0&&num<60:
	// 	fmt.Println("不及格")
	// default:
	// 	fmt.Println("输入错误")
	// }

	// var month int
	// fmt.Println("请输入月份")
	// fmt.Scanf("%d",&month)
	// switch month{
	// case 12,1,2:
	// 	fmt.Println("冬季")
	// case 3,4,5:
	// 	fmt.Println("春季")
	// case 6,7,8:
	// 	fmt.Println("夏季")
	// case 9,10,11:
	// 	fmt.Println("秋季")
	// default:
	// 	fmt.Println("书尼玛内")
	// }

	var a string
	fmt.Println("请输入星期")
	fmt.Scanf("%s", &a)
	switch a {
	case "星期一":
		fmt.Print("干煸毛豆")
	case "星期二":
		{
			fmt.Println("醋溜土豆")
		}

	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期天", "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("输入错误")
	}
}
