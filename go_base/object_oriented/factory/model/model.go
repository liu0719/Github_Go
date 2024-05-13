// 2023/1/7,13:00
package model

type Student struct {
	Name  string
	Score float64
}

type student100 struct {
	Name  string
	score float64
}

func NewStudent(name string, score float64) *student100 {
	return &student100{
		name,
		score,
	}
}

// GetScore 类似Java的get，提高安全性
func (s *student100) GetScore() float64 {
	return s.score
}
