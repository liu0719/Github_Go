package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//file是一个文件对象，文件指针
	fileName := "E:/go/src/go_study/file/open&close&read/text.html"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//带缓冲的方式
	//创建一个*Reader,是带缓冲的
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到换行就结束,空行也会读取
		fmt.Print(str)
		if err == io.EOF { //io.EOF代表文件末尾
			break
		}

	}
	fmt.Println("\n文件读取结束")

	//及时关闭，否则会有内存泄漏
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 2读取文件方式二，一次性读取适用于文件较小的情况！！！！！！
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("read", err)
	}
	fmt.Println(string(content))
}
