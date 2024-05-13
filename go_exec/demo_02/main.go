package main

import "fmt"

func temperature(t float64) string {
	str := ""
	// fmt.Println(&str)
	// write code here
	defer func(p *string) {
		if err := recover(); err != nil {
			*p = "体温异常"
			fmt.Println("*p", *p)
		}
	}(&str)
	if t > 37.5 {
		panic("体温异常")
	}

	return str
}
func main() {
	str := temperature(38.1000)
	// str := ""
	// p := &str
	// defer func() {
	// 	*p = "买买买"
	// }()

	fmt.Println("str", str)
}
