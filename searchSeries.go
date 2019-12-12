package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func searchSerie(serie, param string) (series, int) {
	var search series
	query := url.QueryEscape(serie)
	link := "https://api.thetvdb.com/search/series?" + param + "=" + query

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Authorization", authorization)

	// Execute the request and store the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &search)
	if err != nil {
		panic(err)
	}

	return search, resp.StatusCode
}

type series struct {
	Data []struct {
		Aliases    []string `json:"aliases"`
		Banner     string   `json:"banner"`
		FirstAired string   `json:"firstAired"`
		ID         int      `json:"id"`
		Network    string   `json:"network"`
		Overview   string   `json:"overview"`
		SeriesName string   `json:"seriesName"`
		Slug       string   `json:"slug"`
		Status     string   `json:"status"`
	} `json:"data"`
}
