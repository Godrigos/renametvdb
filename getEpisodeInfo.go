package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getEpisodeInfo(serieName, num string) (episodes, int) {
	var eps episodes
	serie, code := searchSerie(serieName, "name")
	var newCode int

	if code == 401 {
		panic("Not authorized!")
	} else if code == 404 {
		newCode = 405
	} else {
		link := "https://api.thetvdb.com/series/" +
			fmt.Sprintf("%v", serie.Data[0].ID) +
			"/episodes/query?absoluteNumber=" + num

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

		err = json.Unmarshal(body, &eps)
		if err != nil {
			panic(err)
		}
		newCode = resp.StatusCode
	}

	return eps, newCode
}

type episodes struct {
	Data []struct {
		AbsoluteNumber     int      `json:"absoluteNumber"`
		AiredEpisodeNumber int      `json:"airedEpisodeNumber"`
		AiredSeason        int      `json:"airedSeason"`
		AirsAfterSeason    int      `json:"airsAfterSeason"`
		AirsBeforeEpisode  int      `json:"airsBeforeEpisode"`
		AirsBeforeSeason   int      `json:"airsBeforeSeason"`
		Director           string   `json:"director"`
		Directors          []string `json:"directors"`
		DvdChapter         int      `json:"dvdChapter"`
		DvdDiscid          string   `json:"dvdDiscid"`
		DvdEpisodeNumber   int      `json:"dvdEpisodeNumber"`
		DvdSeason          int      `json:"dvdSeason"`
		EpisodeName        string   `json:"episodeName"`
		Filename           string   `json:"filename"`
		FirstAired         string   `json:"firstAired"`
		GuestStars         []string `json:"guestStars"`
		ID                 int      `json:"id"`
		ImdbID             string   `json:"imdbId"`
		LastUpdated        int      `json:"lastUpdated"`
		LastUpdatedBy      int      `json:"lastUpdatedBy"`
		Overview           string   `json:"overview"`
		ProductionCode     string   `json:"productionCode"`
		SeriesID           int      `json:"seriesId"`
		ShowURL            string   `json:"showUrl"`
		SiteRating         float32  `json:"siteRating"`
		SiteRatingCount    int      `json:"siteRatingCount"`
		ThumbAdded         string   `json:"thumbAdded"`
		ThumbAuthor        int      `json:"thumbAuthor"`
		ThumbHeight        string   `json:"thumbHeight"`
		ThumbWidth         string   `json:"thumbWidth"`
		Writers            []string `json:"writers"`
	} `json:"data"`
	Errors struct {
		InvalidFilters     []string `json:"invalidFilters"`
		InvalidLanguage    string   `json:"invalidLanguage"`
		InvalidQueryParams []string `json:"invalidQueryParams"`
	} `json:"errors"`
	Links struct {
		First    int `json:"first"`
		Last     int `json:"last"`
		Next     int `json:"next"`
		Previous int `json:"previous"`
	} `json:"links"`
}
