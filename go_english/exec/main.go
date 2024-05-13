package main

import (
	"english/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	word := &model.Word{
		English: "name",
		Chinese: "名字",
	}
	url := "https://fanyi.so.com/audio?from=en&to=zh&voice=5&cate=uk-speech&key=5d41402abc4b2a76b9719d911017c592&query=" + word.English
	GetPronounceByInter(word, url)
}
func GetPronounceByInter(word *model.Word, url string) error {

	res, err := http.Get(url)
	if err != nil {
		if errs := recover(); err != nil {
			fmt.Println("网络获取发音出错:", errs)
			fmt.Println("请检查网络状况...")
			time.Sleep(time.Second * 3)
			return err
		}

	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取获取发音出错:", err)
		return err
	}
	fmt.Println(data)
	file, err := os.OpenFile(word.Chinese+".mp3", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(data)
	return nil
}
