package main

//goto
func numberOfStepsGoTo(num int) int {
	var step int = 0
	if num == 0 {
		return step
	}
fo:
	step += 1
	if num%2 == 0 {
		num = num / 2
	} else {
		num -= 1
	}
	if num != 0 {
		goto fo
	}
	return step

}

//for
func numberOfStepsFor(num int) int {
	var step int = 0
	for num != 0 {
		step += 1
		if num%2 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
	}

	return step

}

func findRestaurant(list1 []string, list2 []string) []string {

	res := make([]string, 0)
	index := 0
	for i, v := range list1 {

		for j, x := range list2 {

			if x == v {

				if len(res) == 0 {
					res = append(res, x)
					index = i + j
					break

				}
				if index > i+j {
					res = append(res[:len(res)-1], x)
					index = i + j
					break
				}
				if index == i+j {
					res = append(res, x)
					break
				}

			}
		}

	}
	return res
}

func romanToIntMy(s string) int {
	arr1 := []int{1, 5, 10, 50, 100, 500, 1000}
	arr2 := []string{"I", "V", "X", "L", "C", "D", "M"}
	num := 0
	s1 := ToArray(s)
	for i := 0; i < len(s1); i++ {
	lable1:
		for j := range arr2 {
			if s1[i] == arr2[j] {
				if s1[i] == "I" && i < len(s1)-1 && s1[i+1] == "V" {
					num += 4
					i += 1
					break lable1
				} else if s1[i] == "I" && i < len(s1)-1 && s1[i+1] == "X" {
					num += 9
					i += 1
					break lable1
				} else if s1[i] == "X" && i < len(s1)-1 && s1[i+1] == "L" {
					num += 40
					i += 1
					break lable1
				} else if s1[i] == "X" && i < len(s1)-1 && s1[i+1] == "C" {
					num += 90
					i += 1
					break lable1
				} else if s1[i] == "C" && i < len(s1)-1 && s1[i+1] == "D" {
					num += 400
					i += 1
					break lable1
				} else if s1[i] == "C" && i < len(s1)-1 && s1[i+1] == "M" {
					num += 900
					i += 1
					break lable1
				} else {
					num += arr1[j]
				}
			}
		}
	}
	return num
}
func ToArray(s string) []string {
	arr := []string{}
	for _, v := range s {
		arr = append(arr, string(v))
	}
	return arr
}

func getSum(i string) int {
	switch i {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	}
	return 1000
}

//func romanToInt(s string) int {
//	num := 0
//	firstv := getSum(string(s[0]))
//	for i, v := range s {
//		n := getSum(string(s[i]))
//		if n < firstv {
//
//		}
//	}
//}
func main() {
	//fmt.Println(romanToInt("III"))
	//[]
	//	[]
	//list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	//list2 := []string{"KFC", "Shogun", "Burger King"}
	//fmt.Println(findRestaurant(list1, list2))
	//fmt.Println(numberOfStepsFor(14))
	//fmt.Println(100 & 1)
}
