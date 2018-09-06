package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/konojunya/musi/model"
)

func GetTracks(token string) (*model.PlayList, error) {
	values := url.Values{}
	values.Add("q", "hello")
	values.Add("type", "playlist")
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playlist *model.PlayList
	json.Unmarshal(b, &playlist)

	return playlist, nil
}
