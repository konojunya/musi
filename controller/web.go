package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/musi/service"
)

func Index(c *gin.Context) {
	cookie, err := c.Cookie("musi-token")
	if (err != nil) || (cookie == "") {
		c.HTML(http.StatusInternalServerError, "first.html", nil)
		return
	}

	err = service.GetTracks(cookie)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "first.html", nil)
		return
	}

	c.HTML(http.StatusOK, "index.html", nil)
}
