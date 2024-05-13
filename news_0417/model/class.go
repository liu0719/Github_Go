package model

type Class struct {
	Id    int
	Class string
}

func (Class) TableName() string {
	return "class"
}
