package providers

import (
	"encoding/base64"
	"encoding/json"
)

// SearchPayload ...
type SearchPayload struct {
	Method       string      `json:"method"`
	CallbackURL  string      `json:"callback_url"`
	SearchObject interface{} `json:"search_object"`
}

// MovieSearchObject ...
type MovieSearchObject struct {
	IMDBId string            `json:"imdb_id"`
	Title  string            `json:"title"`
	Year   int               `json:"year"`
	Titles map[string]string `json:"titles"`
}

// SeasonSearchObject ...
type SeasonSearchObject struct {
	IMDBId string            `json:"imdb_id"`
	TVDBId int               `json:"tvdb_id"`
	Title  string            `json:"title"`
	Season int               `json:"season"`
	Year   int               `json:"year"`
	Titles map[string]string `json:"titles"`
}

// EpisodeSearchObject ...
type EpisodeSearchObject struct {
	IMDBId         string            `json:"imdb_id"`
	TVDBId         int               `json:"tvdb_id"`
	Title          string            `json:"title"`
	Season         int               `json:"season"`
	Episode        int               `json:"episode"`
	Year           int               `json:"year"`
	Titles         map[string]string `json:"titles"`
	AbsoluteNumber int               `json:"absolute_number"`
}

func (sp *SearchPayload) String() string {
	b, err := json.Marshal(sp)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
