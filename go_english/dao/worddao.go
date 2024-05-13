package dao

import (
	"english/model"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type WordDao struct {
}

func (w WordDao) GetWordList() (*model.WordList, error) {
	file, err := os.OpenFile("./资料.txt", os.O_RDONLY, 0666)
	if err != nil {
		w.Error()
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		w.Error()
	}
	strData := string(data)
	result := strings.Split(strData, "\r\n")

	wordlist := model.WordList{}
	for i := 0; i < len(result); i++ {
		if result[i] == "" {
			result = append(result[:i], result[i+1:]...)
			continue
		}
		word := &model.Word{}
		if i%2 == 0 {
			word.English = result[i]
		}
		i++
		word.Chinese = result[i]

		wordlist.Words = append(wordlist.Words, word)
		wordlist.Date = time.Now().Format("2006-01-02")
	}

	return &wordlist, nil
}
func (w WordDao) Error() {
	fmt.Println("请替代有效的资料.txt,不能为空！！")
	fmt.Println("此窗口3秒后关闭...")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
