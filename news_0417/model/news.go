package model

type News struct {
	Id         int    //新闻ID
	Title      string //新闻标题
	Author     string //作者
	Context    string //内容
	ClassId    int    //所属类型
	CreateTime string //添加时间
	Class      Class  `gorm:"foreignKey:ClassId"`
}

func (News) TableName() string {
	return "news_0417"
}
