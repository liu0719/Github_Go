package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 1.创建一个新文件，写入5句内容"hello Gardon"
	// 1.打开文件
	filePath := "E:/go/src/go_study/file/write/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	// 打开直接defer关闭
	defer file.Close()

	//准备写入5句话
	writer := bufio.NewWriter(file)
	str := "hello Gardon\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// writer是代缓存的,因此调用方法writerString时
	//其实内容是先写到缓存的,所以需要调用flush方法
	//将缓冲的数据真正的写入到文件当中，否则文件中会没有数据

	writer.Flush()

}
