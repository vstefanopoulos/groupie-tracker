package handlers

import "net/http"

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/templates/artists.html") // Serve static HTML
}
