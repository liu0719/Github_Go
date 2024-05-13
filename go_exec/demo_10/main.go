package main

import (
	"fmt"
	"go_exec/demo_10/model"
)

func main() {
	user1 := &model.User{
		Name:    "张三",
		Address: "背景",
		Tele:    "1678189389",
	}
	team1 := &model.Team{
		Name:    "开心团队",
		Address: "天津",
		User:    user1,
	}
	fmt.Println(team1.User.Address)
}
func GetMax[A int | float64](a, b A) A {
	if float64(a) >= float64(b) {
		return a
	}
	return b
}

type int128 int
