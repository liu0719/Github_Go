package utils

import "fmt"

var (
	Age  int
	Name string
)

func init() {
	fmt.Println("utils包里的init")
	Age = 100
	Name = "haha"
}
