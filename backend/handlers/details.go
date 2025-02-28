package handlers

import (
	"fmt"
	"gp/backend/db"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
		"format": formatLocation,
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

func formatLocation(input string) string {
	cleaned := strings.ReplaceAll(input, "_", " ")
	cleaned = strings.ReplaceAll(cleaned, "-", ", ")
	words := strings.Fields(cleaned)

	for i, word := range words {
		switch strings.ToLower(word) {
		case "usa":
			words[i] = "USA"
		case "uae":
			words[i] = "UAE"
		default:
			if len(word) > 0 {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
	}

	return strings.Join(words, " ")
}
