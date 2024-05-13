package main

import (
	"bufio" //读写文件
	"fmt"
	"io"
	"os" //打开关闭文件
)

func main() {
	filePath := "E:/go/src/go_study/file/writer_04/main_04.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}

	//只适用小文件
	content, _ := os.ReadFile(filePath)
	fmt.Println(string(content))

	fmt.Println("大文件——————————————————————————————————————————")
	//大文件件读取方式
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	str := "你好 GoLang\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
