package api

import (
	"encoding/json"
	"gp/backend/db"
	"gp/backend/finder/search"
	"net/http"
	"strings"
)

func SearchAPI(w http.ResponseWriter, r *http.Request) {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	db := db.AllArtists
	var feed []*search.SearchItem = []*search.SearchItem{}
	if r.Method == http.MethodGet {
		searchQuery := strings.TrimSpace(r.URL.Query().Get("q"))
		if searchQuery != "" {
			feed = search.Search(searchQuery, db)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
	return
}
