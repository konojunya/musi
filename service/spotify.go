package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetTracks(token string) error {
	values := url.Values{}
	values.Add("q", "hello")
	values.Add("type", "playlist")
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
