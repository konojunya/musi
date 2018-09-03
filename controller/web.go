package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var config oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config = oauth2.Config{
		ClientID:     os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
		RedirectURL: "http://localhost:4000/oauth",
		Scopes:      []string{"playlist-modify", "user-read-private", "user-library-read"},
	}
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, config.AuthCodeURL("spotify"))
}

func Callback(c *gin.Context) {
	code := c.Query("code")

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	c.SetCookie("spotify-token", token.AccessToken, 1000*60*60*7, "/", "http://localhost:4000", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
