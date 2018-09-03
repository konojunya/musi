package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTracks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hoge": "hoge",
	})
}
