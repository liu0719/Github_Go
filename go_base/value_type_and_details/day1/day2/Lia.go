package day2

func twoSum(nums []int, target int) []int {
	var arr=[]int{}
	for i := 0; i < len(nums); i++ {
		for j:= i; j <len(nums); j++ {
			if nums[i]+nums[j] == target{
				arr = append(arr, i)
				arr= append(arr, j)
				break
			}
		}
	}
	return arr
}