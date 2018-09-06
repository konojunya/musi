package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	cookie, err := c.Cookie("musi-token")
	if (err != nil) || (cookie == "") {
		c.HTML(http.StatusInternalServerError, "first.html", nil)
		return
	}

	c.HTML(http.StatusOK, "index.html", nil)
}
