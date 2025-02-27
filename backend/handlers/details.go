package handlers

import (
	"fmt"
	commonfuncs "gp/backend/common/funcs"
	"gp/backend/db"
	"html/template"
	"net/http"
	"strconv"
)

// Calls APIs and serves Artist DetailsHandler page
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	artistIDStr := r.URL.Query().Get("id")
	if artistIDStr == "" {
		http.Error(w, "Artist ID is required", http.StatusBadRequest)
		return
	}

	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	index, err := db.GetIndex(artistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch artist: %v", err), http.StatusNotFound)
		return
	}

	db.Mutex.Lock()
	artist := db.AllArtists[index]
	db.Mutex.Unlock()

	if artist == nil {
		http.Error(w, fmt.Sprintf("artist pointer is nil: %v", err), http.StatusInternalServerError)
		return
	}

	data := struct {
		Artist   *db.Artist
		Relation db.Relation
	}{
		Artist:   artist,
		Relation: *artist.Relation,
	}

	funcMap := template.FuncMap{
		"format": commonfuncs.FormatLocation,
	}

	tmpl, err := template.New("details.html").Funcs(funcMap).ParseFiles("ui/templates/details.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Failed to execute template: %v", err), http.StatusInternalServerError)
	}
}
