package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	owm "github.com/briandowns/openweathermap"

	"github.com/konojunya/musi/model"
)

func GetTracks(token string, location model.GeoLocation) (*model.Response, error) {
	// weather
	w, err := owm.NewCurrent("F", "EN", openweatherAPIKey)
	if err != nil {
		return nil, err
	}
	w.CurrentByCoordinates(&owm.Coordinates{
		Longitude: location.Longitude,
		Latitude:  location.Latitude,
	})

	values := url.Values{}
	values.Add("q", w.Weather[0].Main)
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

	var playlist model.PlayList
	json.Unmarshal(b, &playlist)

	response := &model.Response{
		CityName: w.Name,
		Playlist: playlist,
		Weather:  w.Weather[0].Main,
	}

	return response, nil
}
