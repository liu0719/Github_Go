package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArtcileController struct {
}

func (art ArtcileController) Index(c *gin.Context) {
	c.String(http.StatusOK, "--文章列表--")
}
func (art ArtcileController) Add(c *gin.Context) {
	c.String(http.StatusOK, "--文章列表--add--")
}
func (art ArtcileController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "--文章列表--edit--")
}
