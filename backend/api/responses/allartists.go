package api

import (
	"encoding/json"
	"gp/backend/db"
	"net/http"
)

func ArtistsAPI(w http.ResponseWriter, r *http.Request) {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.AllArtists)
	return
}
