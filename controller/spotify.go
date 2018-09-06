package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/musi/service"
)

func Login(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, service.GetRedirectURL())
}

func OAuth(c *gin.Context) {
	code := c.Query("code")
	token, err := service.GetToken(code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// c.SetCookie("musi-token", token.AccessToken, 1000*60*60*24*7, "/", "http://localhost:4000", false, false)
	c.SetCookie("musi-token", token.AccessToken, 1000*60, "/", "http://localhost:4000", false, false) // debug
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
