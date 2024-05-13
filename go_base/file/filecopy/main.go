package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(fileName string, dstfileName string) (written int64, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件出错啦：", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	//打开dstfileName
	dstfile, err := os.OpenFile(dstfileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer dstfile.Close()
	writer := bufio.NewWriter(dstfile)
	return io.Copy(writer, reader)
}

//将本文件夹里的go导图.png拷贝到桌面上
func main() {

	filePath := "E:/go/src/go_study/file/filecopy/go导图.png"
	copyPath := "C:/Users/33542/Desktop/go_study.png"
	_, err := CopyFile(filePath, copyPath)
	if err != nil {
		fmt.Println("拷贝失败：", err)

	} else {
		fmt.Println("拷贝成功")
	}

}
