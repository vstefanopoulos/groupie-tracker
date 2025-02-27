package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gp/backend/api/calls"
	"gp/backend/db"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestAPIs(t *testing.T) {
	os.Chdir("../..")
	mux := http.NewServeMux()

	// Load data from APi
	err := calls.FetchAll()
	if err != nil {
		log.Fatalf("Error initializing artists: %v", err)
	}

	// Give time for Relations to load
	time.Sleep(1 * time.Second)

	mux.HandleFunc("/api/search", SearchAPI)
	mux.HandleFunc("/api/filter", FiltersAPI)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Test Search
	t.Run("Test SEARCH=Queen", func(t *testing.T) {
		req, err := http.NewRequest("GET", server.URL+"/api/search?q=queen", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []struct {
			Name string `json:"name"`
		}

		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test SEARCH=london-uk", func(t *testing.T) {
		req, err := http.NewRequest("GET", server.URL+"/api/search?q=london-uk", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []struct {
			Name string `json:"name"`
		}

		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test SEARCH=05-08-1967", func(t *testing.T) {
		req, err := http.NewRequest("GET", server.URL+"/api/search?q=05-08-1967", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []struct {
			Name string `json:"name"`
		}

		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test SEARCH=creation 1973", func(t *testing.T) {
		req, err := http.NewRequest("GET", server.URL+"/api/search?q=creation 1973", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []struct {
			Name string `json:"name"`
		}

		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test FILTERS by creation=[1965,1965]", func(t *testing.T) {
		payload := []byte(`{"creation": [1965,1965]}`)
		req, err := http.NewRequest("POST", server.URL+"/api/filter", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("Error creating POST request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []*db.Artist
		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test FILTERS by locations=Texas, USA", func(t *testing.T) {
		payload := []byte(`{"locations": "Texas, USA"}`)
		req, err := http.NewRequest("POST", server.URL+"/api/filter", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("Error creating POST request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req) // Simulate request

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []*db.Artist
		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})

	t.Run("Test FILTERS by {locations: Washington, USA, members: 4}", func(t *testing.T) {
		payload := []byte(`{"locations": "Washington, USA", "members": [4]}`)
		req, err := http.NewRequest("POST", server.URL+"/api/filter", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("Error creating POST request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rec.Code)
		}

		var artists []*db.Artist
		err = json.NewDecoder(rec.Body).Decode(&artists)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		fmt.Println("\nResults:")
		for _, artist := range artists {
			fmt.Printf("Artist: %v\n", artist.Name)
		}
		fmt.Print("\n")
	})
}
