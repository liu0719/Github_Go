package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "E:/go/src/go_study/file/contemt_copy/index.html"
	filePathCopy := "E:/go/src/go_study/file/contemt_copy/copy.html"
	reader, _ := ioutil.ReadFile(filePath)
	fmt.Print(string(reader))
	//会覆盖文件中的内容
	err1 := ioutil.WriteFile(filePathCopy, reader, 0666) //如果没有文件，会自动创建文件
	if err1 != nil {
		fmt.Println(err1)
	}
	// readercopy, _ := ioutil.ReadFile(filePathCopy)
}
