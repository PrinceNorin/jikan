package jikan

type Service interface {
	SearchAnime(query string) ([]*Anime, error)
}

type Anime struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
	MalURL   string `json:"malUrl"`
	ShowType string `json:"showType"`
	Status   string `json:"status"`
}
