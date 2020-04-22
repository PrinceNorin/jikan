package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PrinceNorin/jikan"
)

const apiURL = "https://api.jikan.moe/v3/search/anime?limit=16&q=%s"

type jikanResult struct {
	Results []*jikanAnime `json:"results"`
}

type jikanAnime struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Airing   bool   `json:"airing"`
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
}

func New() jikan.Service {
	return &service{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type service struct {
	http *http.Client
}

func (s *service) SearchAnime(query string) ([]*jikan.Anime, error) {
	res, err := s.http.Get(fmt.Sprintf(apiURL, url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r jikanResult
	if err := json.Unmarshal(buf, &r); err != nil {
		return nil, err
	}

	animes := make([]*jikan.Anime, 1)
	for _, ja := range r.Results {
		status := "Finished"
		if ja.Airing {
			status = "Ongoing"
		}

		animes = append(animes, &jikan.Anime{
			MalURL:   ja.URL,
			ImageURL: ja.ImageURL,
			Title:    ja.Title,
			ShowType: ja.Type,
			Status:   status,
		})
	}

	return animes, nil
}
