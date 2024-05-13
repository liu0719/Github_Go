package main

import (
	"fmt"
	"sort"
)

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}

func isSameAfterReversals(num int) bool {
	return num%10 > 0
}

func makeGoodMy(s string) string {
	flag := 1
	arr := []byte(s)
	for {
		if flag >= len(arr) {
			break
		}
		var char = int(arr[flag]) - int(arr[flag-1])
		if char == 32 || char == -32 {
			arr = append(arr[:flag-1], arr[flag+1:]...)
			flag = 1
			continue
		}
		flag++
	}
	return string(arr)
}

func makeGood(s string) string {

	arr := []byte(s)
	for flag := 1; flag < len(arr); flag += 1 {
		var char = int(arr[flag]) - int(arr[flag-1])
		if char == 32 || char == -32 {
			arr = append(arr[:flag-1], arr[flag+1:]...)
			flag = 0
		}
	}
	return string(arr)
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	var res int
	for _, v := range arr {
		if res < v {
			res++
		}
	}
	return res
}

func trap(height []int) int {
	res := 0
	n := len(height)
	if n <= 2 {
		return res
	}

	l, r := 0, n-1
	lH, rH := height[l], height[r]
	for l+1 < r {
		if lH < rH {
			if lH > height[l+1] {
				res += lH - height[l+1]
			} else {
				lH = height[l+1]
			}
			l++
		} else {
			if rH > height[r-1] {
				res += rH - height[r-1]
			} else {
				rH = height[r-1]
			}
			r--
		}
	}

	return res
}

func goodDaysToRobBank(security []int, time int) []int {
	if time == 0 {
		return nil
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	a, b := 0, 0
lable:
	for i, v1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target {
				a = i
				b = j
				break lable
			}
		}
	}
	return []int{a, b}
}

func isPalindrome(x int) bool {
	// 倒序后  判断是不是和原来的数字相等
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x != 0 {
		redirect = redirect*10 + x%10
		x /= 10
	}
	return origin == redirect
}
func main() {
	arr1 := []int{3, 3}
	fmt.Println(twoSum(arr1, 6))
	//arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//fmt.Println(makeGood("abBAcC"))
	//fmt.Println(trap(arr))
	//fmt.Println(goodDaysToRobBank(arr, 0))
	//fmt.Println(0%10 > 0)
}
