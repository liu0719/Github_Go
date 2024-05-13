package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type charharcharount struct {
	Eng      int
	Num      int
	Spachare int
	Other    int
	chinese  int
}

func main() {
	filepath := "E:/go/src/go_study/file/exec/main.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	var char charharcharount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		str1 := []rune(str)
		for _, v := range str1 {
			switch {
			case v >= 'A' && v <= 'Z':
				fallthrough
			case v >= 'a' && v <= 'z':
				char.Eng++
			case v >= '0' && v <= '9':
				char.Num++
			case v == ' ' || v == '\t':
				char.Spachare++
			case unicode.Is(unicode.Han, v):
				char.chinese++
			default:
				char.Other++
			}

		}
		if err == io.EOF { //文件末尾退出
			break
		}
	}
	//输出统计的字符
	fmt.Printf("中文:%v,英文:%v,数字:%v,空格:%v,其他:%v", char.chinese, char.Eng, char.Num, char.Spachare, char.Other)
}
