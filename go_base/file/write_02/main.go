package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "E:/go/src/go_study/file/write_02/main.txt"
	//									只写模式，会覆盖文件之前的内容
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//及时关闭
	defer file.Close()
	//覆盖内容
	str := "你好，尚硅谷\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
