package calls

import (
	"encoding/json"
	"fmt"
	"gp/backend/db"
	"net/http"
)

// Fetches data from locations API
func FetchLocation(id int) (db.Location, error) {
	locResp, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id))
	if err != nil {
		return db.Location{}, fmt.Errorf("%w %v", ErrFailedToFetch, err)
	}
	defer locResp.Body.Close()
	var location db.Location

	if err := json.NewDecoder(locResp.Body).Decode(&location); err != nil {
		return db.Location{}, fmt.Errorf("fetch location: %w %v", ErrFailedToDecode, err)
	}

	return location, nil
}
