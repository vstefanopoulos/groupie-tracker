package calls

import (
	"encoding/json"
	"fmt"
	"gp/backend/db"
	"log"
	"net/http"
)

func FetchRelations() (map[int]db.Relation, error) {
	log.Println("Fetching all relations from API...")

	resp, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation"))
	if err != nil {
		return nil, fmt.Errorf("%w %v", ErrFailedToFetch, err)
	}
	defer resp.Body.Close()

	var response struct {
		Index []*db.Relation `json:"index"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("fetch relations: %w %v", ErrFailedToDecode, err)
	}

	relation := make(map[int]db.Relation)
	for _, entry := range response.Index {
		relation[entry.ID] = *entry
	}

	return relation, nil
}
