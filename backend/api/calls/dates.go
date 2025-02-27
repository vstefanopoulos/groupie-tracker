package calls

import (
	"encoding/json"
	"fmt"
	"gp/backend/db"
	"net/http"
)

// Fetches data from dates API
func FetchDate(id int) (db.Date, error) {
	datesResp, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id))
	if err != nil {
		return db.Date{}, fmt.Errorf("%w %v", ErrFailedToFetch, err)
	}
	defer datesResp.Body.Close()
	var date db.Date

	if err := json.NewDecoder(datesResp.Body).Decode(&date); err != nil {
		return db.Date{}, fmt.Errorf("fetch date: %w %v", ErrFailedToDecode, err)
	}

	return date, nil
}
