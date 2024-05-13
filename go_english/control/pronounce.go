package control

import (
	"english/dao"
	"english/model"
	"english/utils"
	"fmt"
	"time"
)

type PronounceControl struct {
}

func (p PronounceControl) OverAllControl(word *model.Word) {
	proDao := &dao.PronounceDao{}

	if err := proDao.GetPronounceByLocal(word); err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		err = proDao.GetPronounceByInter(word)
		if err != nil {
			utils.ShowMessage("获取单词发音失败", "网络错误或者文件格式不支持")
		}
		err = proDao.GetPronounceByLocal(word)
		if err != nil {
			utils.ShowMessage("获取单词发音失败", "网络错误或者文件格式不支持")
		}

	}

}
