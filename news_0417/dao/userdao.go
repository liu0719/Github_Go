package dao

import (
	"news_0417/model"
	"news_0417/utils"
)

// 新增新闻
func AddNews(new *model.News) error {
	result := utils.DB.Save(new)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除新闻
func DeleteNew(newId string) error {
	result := utils.DB.Where("id = ?", newId).Delete(model.News{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新新闻
func UpdateNews(new *model.News) error {
	result := utils.DB.Save(new)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 查找所有新闻
func GetNews() ([]model.News, error) {
	var news []model.News
	result := utils.DB.Model(model.News{}).Preload("Class").Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

// 根据class查找新闻
func GetNewsByClass(classId string) ([]model.News, error) {
	var news []model.News
	result := utils.DB.Where("class_id=?", classId).Preload("Class").Model(model.News{}).Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

// 根据Id查找新闻
func GetNewsById(Id string) (*model.News, error) {
	var news model.News
	result := utils.DB.Where("id=?", Id).Preload("Class").Model(model.News{}).First(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return &news, nil
}

// 根据作者查找新闻
func GetNewsByAuthor(a string) ([]model.News, error) {
	var news []model.News
	result := utils.DB.Where("author = ?", a).Preload("Class").Model(model.News{}).Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

// 根据作者和classId查找新闻
func GetNewsByAuthorAndClass(a string, classId string) ([]model.News, error) {
	var news []model.News
	result := utils.DB.Where("author = ? and class_id=?", a, classId).Preload("Class").Model(model.News{}).Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}
