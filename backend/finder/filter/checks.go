package filter

import (
	"gp/backend/db"
	"gp/backend/finder/shared"
	"strconv"
)

// Range over artists locations and return true if any location matches the query
func byLocations(query string, artist *db.Artist) bool {
	if query == "" {
		return true
	}

	if artist.Relation == nil {
		return false
	}

	if checkRelations(shared.Normalize(query), artist.Relation.DatesLocations) {
		return true
	}

	return false
}

func byFirstAlbum(firstAlbum []int, artist *db.Artist) bool {
	if len(firstAlbum) == 0 {
		return true
	}
	firstAlbumInt, _ := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
	return firstAlbumInt >= firstAlbum[0] && firstAlbumInt <= firstAlbum[1]
}

func byMembers(members []int, artist *db.Artist) bool {
	if len(members) == 0 {
		return true
	}
	for _, num := range members {
		if len(artist.Members) == num {
			return true
		}
	}
	return false
}

func byCreation(searchDates []int, artist *db.Artist) bool {
	if len(searchDates) < 2 {
		return true
	}
	return artist.CreationDate >= searchDates[0] && artist.CreationDate <= searchDates[1]
}

func checkRelations(query string, locations map[string][]string) bool {
	for location := range locations {
		if shared.IsContained(query, shared.Normalize(location)) {
			return true
		}
	}
	return false
}
