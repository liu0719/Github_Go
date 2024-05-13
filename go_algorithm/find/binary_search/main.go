// 2023/1/3,10:35
package main

import "fmt"

//二分查找只能在有序的数组中进行
//二分查找for循环
func binaryFind(list []int, num int) (index int) {
	index = -1
	lIndex, rIndex := 0, len(list)-1
	for {
		mIndex := (lIndex + rIndex) / 2
		if list[mIndex] > num {
			rIndex = mIndex - 1
		}
		if list[mIndex] < num {
			lIndex = mIndex + 1
		}
		if list[mIndex] == num {
			index = mIndex
			break
		}
		if lIndex > rIndex {
			break
		}
	}
	return index
}

//二分查找递归

func binaryFindRecursion(arr []int, leftInedx int, rightIndex, num int) {
	if leftInedx > rightIndex {
		fmt.Println("没找到")
	}
	middleIndex := (leftInedx + rightIndex) / 2
	if arr[middleIndex] > num {
		binaryFindRecursion(arr, leftInedx, middleIndex-1, num)
	}
	if arr[middleIndex] < num {
		binaryFindRecursion(arr, middleIndex+1, rightIndex, num)
	}
	if arr[middleIndex] == num {
		fmt.Println("找到了霞表示：", middleIndex)
	}
}

func main() {
	arr := []int{1, 8, 10, 99, 1000, 1234}
	binaryFindRecursion(arr, 0, len(arr)-1, 99)
}
