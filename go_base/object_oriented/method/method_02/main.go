// 2023/1/6,18:24
package main

import "fmt"

type MethodUtils struct {
}

func (met MethodUtils) printMN(m int, n int) {
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) print() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (met MethodUtils) area(w float64, h float64) float64 {
	return h * w
}
func (met MethodUtils) chengfabiao(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v x %v = %v\t", j+1, i+1, (j+1)*(i+1))
		}
		fmt.Println()
	}
}

func (met MethodUtils) 矩阵转换(a [3][3]int) [3][3]int {
	b := [3][3]int{}
	for i, v1 := range a {
		for j := range v1 {
			b[i][j] = a[j][i]
		}
	}
	return b
}
func main() {
	m := MethodUtils{}
	m.printMN(10, 2)
	fmt.Println(m.area(10.234, 23.768))
	m.chengfabiao(9)
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	q := (&m).矩阵转换(arr)
	for _, v := range q {
		fmt.Println(v)
	}
}
