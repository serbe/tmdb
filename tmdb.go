package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiURL = "http://api.themoviedb.org/3/"
)

// TMDB values:
// apikey store your API key from themoviedb
type TMDB struct {
	apiKey string
}

// response of search/multi
type tmdbResponse struct {
	Page          int
	Results       []tmdbResult
	Total_pages   int
	Total_results int
}

// results format from Tmdb
type tmdbResult struct {
	Name              string
	Adult             bool
	Backdrop_path     string
	Genre_ids         []int64
	Id                int64
	Original_language string
	Original_title    string
	Overview          string
	Release_date      string
	Poster_path       string
	Popularity        float32
	Title             string
	Vote_average      float32
	Vote_count        int64
}

// Init tmbd to set API value
func Init(apiKey string) *TMDB {
	return &TMDB{apiKey: apiKey}
}

// GetByName get data from themoviedb by name
func (tmdb *TMDB) GetByName(movieName string) (string, error) {
	var response = &tmdbResponse{}
	resp, err := http.Get(apiURL + "search/movie?api_key=" + tmdb.apiKey + "&language=ru&query=" + url.QueryEscape(movieName))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errorStatus(resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &response)
	if err == nil {
		fmt.Println(response)
	}
	return string(body), err
}

func errorStatus(status int) error {
	return fmt.Errorf("Status Code %d received from TMDb", status)
}
