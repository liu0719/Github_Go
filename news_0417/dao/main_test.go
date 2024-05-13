package dao

import (
	"fmt"
	"news_0417/model"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	// t.Run("添加新闻", testAddNews)
	// t.Run("添加新闻", testUpdateNews)
	// t.Run("删除新闻", testDeleteNews)
	t.Run("删除新闻", testGetNews)
	// t.Run("删除新闻", testGetNewsByClass)
}
func testAddNews(t *testing.T) {
	new := &model.News{
		Id:         1,
		Title:      "新闻1",
		Author:     "张三",
		Context:    "张三的新闻内容",
		ClassId:    3,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := AddNews(new)
	if err != nil {
		fmt.Println(err)
	}
}
func testUpdateNews(t *testing.T) {
	new := &model.News{
		Id:         1,
		Title:      "新闻2",
		Author:     "张三3333",
		Context:    "张三的新闻内容11111",
		ClassId:    1,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := UpdateNews(new)
	if err != nil {
		fmt.Println(err)
	}
}
func testDeleteNews(t *testing.T) {
	err := DeleteNew("1")
	if err != nil {
		fmt.Println(err)
	}
}
func testGetNews(t *testing.T) {
	news, err := GetNews()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range news {
		fmt.Println(v)
	}
}
func testGetNewsByClass(t *testing.T) {
	news, err := GetNewsByClass("3")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range news {
		fmt.Println(v)
	}
}
