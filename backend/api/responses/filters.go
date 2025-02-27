package api

import (
	"encoding/json"
	"gp/backend/db"
	"gp/backend/finder/filter"
	"net/http"
)

func FiltersAPI(w http.ResponseWriter, r *http.Request) {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()
	feed := db.AllArtists
	if r.Method == http.MethodPost {
		filterValues := &filter.Values{}
		if err := json.NewDecoder(r.Body).Decode(filterValues); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		feed = filter.AllFilters(filterValues, feed)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
	return
}
