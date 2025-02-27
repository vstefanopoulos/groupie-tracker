package db

import "fmt"

// Iterates through allArtists and returns index of the artist id. Concurent safe
func GetIndex(id int) (int, error) {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for i, artist := range AllArtists {
		if artist == nil {
			continue
		}
		if artist.ID == id {
			return i, nil
		}
	}

	return 0, fmt.Errorf("get index: %w id: %d", ErrNotInDB, id)
}
