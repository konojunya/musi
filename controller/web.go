package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	cookie, _ := c.Cookie("musi-token")
	if cookie == "" {
		c.HTML(http.StatusOK, "first.html", nil)
		return
	}

	c.HTML(http.StatusOK, "index.html", nil)
}
