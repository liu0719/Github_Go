// 2023/1/1,20:00
package main

import (
	"fmt"
	"os"
)

func eee(a ...any) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}
func main() {
	//数组使用细节和注意
	eee("haha")
}
