package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for i, v := range nums {
		if v != val {
			nums[i] = v
			left++
		}
	}
	return left
}
func main() {
	f := []int{2, 2, 3, 3, 4}
	i := removeElement(f, 3)
	fmt.Println(i)

}
