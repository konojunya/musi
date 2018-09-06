package service

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var (
	openweatherAPIKey string
	config            oauth2.Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	openweatherAPIKey = os.Getenv("openweather")

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

func GetRedirectURL() string {
	// TODO: CSRF対策
	return config.AuthCodeURL("state")
}

func GetToken(code string) (*oauth2.Token, error) {
	return config.Exchange(oauth2.NoContext, code)
}
