package control

import (
	"english/dao"
	"english/model"
	"fmt"
)

type WordControl struct {
}

func (w WordControl) GetWord() *model.WordList {
	worddao := dao.WordDao{}
	wordlist, err := worddao.GetWordList()
	if err != nil {
		fmt.Println("err:", err)
	}
	return wordlist

}
