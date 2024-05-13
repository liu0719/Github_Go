package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

type Person struct {
	Name string
}

func (p *Person) Shuai() string {
	return p.Name + "真帅"
}

func Index(c *gin.Context) {
	p := &Person{
		Name: "刘广硕",
	}
	t := template.Must(template.ParseFiles("static/music/index.html"))
	t.Execute(c.Writer, p)
}
func main() {
	r := gin.Default()
	r.Use(favicon.New("favicon.ico"))
	r.Static("/static/", "./static/")
	r.GET("/", Index)
	r.Run("10.50.142.216:8888")
}
