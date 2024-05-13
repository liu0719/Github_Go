package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	// "io"
	"os"
	// "strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.县创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	//2.输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

	//转成稀疏数组
	//思路
	// 1.遍历chessMap,如果我们发现有一个元素，的值不为零，创建一个node1结构体
	// 2.将其放入对应切边中

	var sparseArr []ValNode

	//创建一个Node结点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	//标准的稀疏数组有一个记录元素的二维数组的规模（行和列，默认值）
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	for i, v := range sparseArr {
		fmt.Println(i, v.row, v.col, v.val)
	}
	//存盘
	filePath := "E:/go/src/go_algorithm/sparse_matrix/spareArr.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)

	for _, v := range sparseArr {
		writer.WriteString(fmt.Sprintf("%v %v %v\n", v.row, v.col, v.val))
	}
	writer.Flush()
	file.Close()
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//读盘
	reader := bufio.NewReader(file)
	var sparseArr2 []ValNode
	for i := 0; i < 3; i++ {
		str, err := reader.ReadString('\n')
		str = strings.Replace(str, "\n", " ", -1)
		strslice := strings.Split(str, " ")
		row1, _ := strconv.Atoi(strslice[0])
		col1, _ := strconv.Atoi(strslice[1])
		val1, _ := strconv.Atoi(strslice[2])
		valNode = ValNode{
			row: row1,
			col: col1,
			val: val1,
		}
		sparseArr2 = append(sparseArr2, valNode)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("--------------------------")
	for _, v := range sparseArr2 {
		fmt.Println(v.row, v.col, v.val)
	}

	//先创建一个原始数组
	var chessMap2 [11][11]int

	for _, v := range sparseArr2[1:] {
		chessMap2[v.row][v.col] = v.val
	}
	fmt.Println("=====================")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
}
