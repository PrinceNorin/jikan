package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PrinceNorin/jikan"
)

func NewHTTP(svc jikan.Service) http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/search", searchAnimeHandler(svc))

	return h
}

func searchAnimeHandler(svc jikan.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		animes, err := svc.SearchAnime(r.URL.Query().Get("q"))
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
		} else {
			json.NewEncoder(w).Encode(animes)
		}
	}
}
