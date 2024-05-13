package dao

import (
	"english/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type PronounceDao struct {
}

func (p *PronounceDao) GetPronounceByLocal(word *model.Word) error {
	f, err := os.Open("./pronounce/" + time.Now().Format("2006-01-02") + "/" + word.English + "-" + word.Chinese + ".mp3")
	if err != nil {

		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return nil
}

// 网络获取
func (p *PronounceDao) GetPronounceByInter(word *model.Word) error {
	wordPath := "./pronounce/" + time.Now().Format("2006-01-02") + "/" + word.English + "-" + word.Chinese + ".mp3"
	url := "https://dict.youdao.com/dictvoice?audio=" + word.English + "&le=en"
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
	file, err := os.OpenFile(wordPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		flag := os.IsNotExist(err)
		if flag {
			path := "./pronounce/" + time.Now().Format("2006-01-02")
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				if !os.IsExist(err) {
					fmt.Println("创建文件夹出错:", err)
				}

			}
			file, err = os.OpenFile(wordPath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("打开文件错误:", err)
			}

		} else {
			file, err = os.OpenFile(wordPath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("打开文件错误:", err)
			}
		}

	}
	file.Write(data)
	word.Pronounce = "./pronounce/" + time.Now().Format("2006-01-02") + "/" + word.English + "-" + word.Chinese + ".mp3"
	defer file.Close()
	return nil
}
