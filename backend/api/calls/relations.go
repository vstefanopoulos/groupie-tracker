package calls

import (
	"encoding/json"
	"fmt"
	"gp/backend/db"
	"net/http"
)

// Fetches data from relations API for Artist with ID=id
func FetchRelation(id int) (*db.Relation, error) {
	resp, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return nil, fmt.Errorf("%w %v", ErrFailedToFetch, err)
	}
	defer resp.Body.Close()

	var relation db.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relation); err != nil {
		return nil, fmt.Errorf("fetch relation: %w %v", ErrFailedToDecode, err)
	}

	return &relation, nil
}
