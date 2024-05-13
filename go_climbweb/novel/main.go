package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	j := 1
	for i := 152999543; i <= 181449187; i++ {
		strI := fmt.Sprintf("%v", i)
		// https://www.xxbiqudu.com/91_91674/153304570.html
		url := "https://www.xxbiqudu.com/91_91674/" + strI + ".html"
		data := GetHtmlByUrl(url)
		content := Getnovel(data)
		if content == "" {
			fmt.Println("出现错误,第", j, "章")
		}
		if !strings.HasSuffix(content, "\n\r") {
			content = content + "\n\r"
		}
		err := WriteToFile(content)
		if err != nil {
			fmt.Println("出现错误", err)
		}
		j++
	}
	fmt.Println("全部写入完毕")

}

// 获取网址的html
func GetHtmlByUrl(url string) string {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	data, err = GbkToUtf8(data)
	if err != nil {
		fmt.Println("err:", err)
	}
	return string(data)
}

// 获取一章的标题和内容
func Getnovel(data string) string {
	title := FindTitle(data)
	content := FindContent(data)
	allContent := title + "\n\r" + content
	if len(allContent) < 10 {
		return ""
	}
	return allContent
}

// 打开文件进行写入
func WriteToFile(data string) error {
	file, err := os.OpenFile("茅山鬼王.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("写入文件错误：", err)
		return nil
	}
	return nil
}

// 解析Html,查找标题
func FindTitle(data string) string {
	var title string
	root, err := htmlquery.Parse(strings.NewReader(data))
	if err != nil {
		fmt.Println("err", err)
	}
	find := htmlquery.Find(root, "//div[@class='bookname']/h1")
	for _, v := range find {
		title = htmlquery.InnerText(v)
	}
	return title
}

// 解析Html,查找文章
func FindContent(data string) string {
	var content string
	root, err := htmlquery.Parse(strings.NewReader(data))
	if err != nil {
		fmt.Println("err", err)
	}
	find := htmlquery.Find(root, "//*[@id='content']")
	for _, v := range find {
		content = htmlquery.InnerText(v)
	}
	return content
}

// 格式转换 GbkToUtf8
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}
