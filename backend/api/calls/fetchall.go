package calls

import (
	"encoding/json"
	"fmt"
	"gp/backend/db"
	"log"
	"net/http"
)

// Populates allArtist var from artists API. Concurrent safe
func FetchAll() error {
	log.Println("Donwloading artists from API...")
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return fmt.Errorf("%w %v", ErrFailedToFetch, err)
	}
	defer resp.Body.Close()

	var artistsTemp []*db.Artist

	if err := json.NewDecoder(resp.Body).Decode(&artistsTemp); err != nil {
		return fmt.Errorf("%w %v", ErrFailedToDecode, err)
	}

	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	db.AllArtists = artistsTemp
	go func() {
		err = fetchRelations()
	}()
	return err
}

// Updates AllArtist instances with Relation
func fetchRelations() (err error) {
	relations, err := FetchRelations()
	if err != nil {
		fmt.Println(err)
		return err
	}

	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	for _, artist := range db.AllArtists {
		if artist == nil {
			continue
		}
		if relation, exists := relations[artist.ID]; exists {
			artist.Relation = &relation
		}
	}

	log.Println("Relations updated to database")
	return nil
}
