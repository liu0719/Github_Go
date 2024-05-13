package main

import "fmt"

func SetWay(mymap *[8][7]int, i int, j int) bool {
	if mymap[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if mymap[i][j] == 0 { //如果这个点可以探测
			//假设是通的，需要探测,上下左右
			//换成下右上左
			mymap[i][j] = 2
			if SetWay(mymap, i+1, j) { //下
				return true
			} else if SetWay(mymap, i, j+1) { //向右
				return true
			} else if SetWay(mymap, i-1, j) { //上
				return true
			} else if SetWay(mymap, i, j-1) { //左
				return true
			} else { //死路
				mymap[i][j] = 3
			}
		} else { //这个点不能探测，1，是墙
			return false
		}
	}
	return false
}

func main() {
	//创建二维数组，模拟迷宫
	//1 就是墙
	//0 就是没有走过的点
	//2 就是一个通路
	//3 走过，但是不通
	var mymap [8][7]int
	for i := 0; i < 7; i++ {
		mymap[0][i] = 1
		mymap[7][i] = 1
	}
	for i := 0; i < 7; i++ {
		mymap[i][0] = 1
		mymap[i][6] = 1
	}
	mymap[3][1] = 1
	mymap[3][2] = 1
	mymap[1][2] = 1
	mymap[2][2] = 1

	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mymap[i][j], " ")
		}
		fmt.Println()
	}
	SetWay(&mymap, 1, 1)
	fmt.Println("探测完成...")
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mymap[i][j], " ")
		}
		fmt.Println()
	}

}
