package main

import "fmt"

func main() {
	//a := 0
	//for i := 0; i < 100; i++ {
	//	a += i
	//	if a > 20 {
	//		fmt.Println(i)
	//		break
	//	}
	//}
	var username string
	var password int
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if username == "张无忌" && password == 888 {

			fmt.Println("登陆成功")
			break
		}
		if 3-1-i == 0 {
			fmt.Println("给你机会你不中用啊")
		} else {
			fmt.Println("输入错误您还有", 3-i-1, "次机会，请重试")

		}
	}

}
