package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "E:/go/src/go_study/file/writer_03/main.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("oepn file err:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "我是后来加上的\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
