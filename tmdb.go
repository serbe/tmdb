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

//{"page":1,"results":[{"adult":false,"backdrop_path":"/xBKGJQsAIeweesB79KC89FpBrVr.jpg","genre_ids":[18,80],"id":278,"original_language":"en","original_title":"The Shawshank Redemption","overview":"Фильм удостоен шести номинаций на `Оскар`, в том числе и как лучший фильм года. Шоушенк - название тюрьмы. И если тебе нет еще 30-ти, а ты получаешь пожизненное, то приготовься к худшему: для тебя выхода из Шоушенка не будет! Актриса Рита Хэйворт - любимица всей Америки. Энди Дифрейну она тоже очень нравилась. Рита никогда не слышала о существовании Энди, однако жизнь Дифрейну, бывшему вице-президенту крупного банка, осужденному за убийство жены и ее любовника, Рита Хэйворт все-таки спасла.","release_date":"1994-09-14","poster_path":"/sRBNv6399ZpCE4RrM8tRsDLSsaG.jpg","popularity":3.836485,"title":"Побег из Шоушенка","video":false,"vote_average":8.2,"vote_count":4154}],"total_pages":1,"total_results":1} <nil>

// results format from Tmdb
type tmdbResult struct {
	Adult          bool
	Name           string
	Backdrop_path  string
	Genre_ids      []int64
	Id             int64
	Original_name  string
	Original_title string
	Overview       string
	Release_date   string
	Poster_path    string
	Popularity     float32
	Title          string
	Vote_average   float32
	Vote_count     int64
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
