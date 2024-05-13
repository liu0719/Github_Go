package models

type Book struct {
	Id      int
	Title   string
	Author  string
	Price   float64
	Sales   int
	Stock   int
	ImgPath string
}

func (Book) TableName() string {
	return "books"
}
