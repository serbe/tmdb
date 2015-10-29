package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	apiURL = "http://api.themoviedb.org/3"
)

// TMDB values:
// apikey store your API key from themoviedb
type TMDB struct {
	apiKey string
	config *tmdbConfig
}

type tmdbConfig struct {
	Images tmdbImagesConfig
}

type tmdbImagesConfig struct {
	Base_url        string
	Secure_base_url string
	Backdrop_sizes  []string
	Logo_sizes      []string
	Poster_sizes    []string
	Profile_sizes   []string
	Still_sizes     []string
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
	Poster_base_url string
}

// Init tmbd to set API value
func Init(apiKey string) *TMDB {
	return &TMDB{apiKey: apiKey}
}

// GetConfig get img configuration from themoviedb
func (tmdb *TMDB) getConfig() error {
	if tmdb.config == nil || tmdb.config.Images.Base_url == "" {
		config := &tmdbConfig{}
		resp, err := http.Get(apiURL + "/configuration?api_key=" + tmdb.apiKey)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("Status Code %d received from TMDb", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = json.Unmarshal(body, &config)
		if err != nil {
			fmt.Println(err)
			return err
		}
		tmdb.config = config
	}
	return nil
}

// GetByName get data from themoviedb by name
func (tmdb *TMDB) GetByName(movieName string, year string) (tmdbResult, error) {
	tmdb.getConfig()
	time.Sleep(1 * time.Second)
	var response = &tmdbResponse{}
	if year != "" {
		year = "&year=" + year
	}
	queryString := apiURL + "/search/movie?api_key=" + tmdb.apiKey + "&language=ru&query=" + url.QueryEscape(movieName) + year
	resp, err := http.Get(queryString)
	if err != nil {
		return tmdbResult{}, err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Header)
		return tmdbResult{}, fmt.Errorf("Status Code %d received from TMDb", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tmdbResult{}, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return tmdbResult{}, err
	}
	if len(response.Results) == 0 {
		return tmdbResult{}, err
	}
	response.Results[0].Poster_base_url = tmdb.config.Images.Base_url
	return response.Results[0], err
}
