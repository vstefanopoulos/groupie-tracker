package handlers

import (
	api "gp/backend/api/responses"
	"net/http"
)

func Handlers() *http.ServeMux {
	groupie := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui/static"))
	groupie.Handle("/static/", http.StripPrefix("/static/", fs))
	groupie.HandleFunc("/", Homepage)
	groupie.HandleFunc("/details", DetailsHandler)
	groupie.HandleFunc("/artists", ArtistsHandler)

	groupie.HandleFunc("/api/artists", api.ArtistsAPI)
	groupie.HandleFunc("/api/search", api.SearchAPI)
	groupie.HandleFunc("/api/filter", api.FiltersAPI)

	return groupie
}
