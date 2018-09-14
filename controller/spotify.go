package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/musi/model"
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

	c.SetCookie("musi-token", token.AccessToken, 1000*60*60*24*7, "/", "https://musi-app.now.sh", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func GetPlayList(c *gin.Context) {
	cookie, _ := c.Cookie("musi-token")
	if cookie == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lng, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	location := model.GeoLocation{
		Longitude: lng,
		Latitude:  lat,
	}

	playlist, err := service.GetTracks(cookie, location)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, playlist)
}
