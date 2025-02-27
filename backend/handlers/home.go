package handlers

import "net/http"

// Serves static home page
func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "UI/templates/index.html")
}
