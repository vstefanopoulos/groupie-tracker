package filter

/*Filter module provides */

import (
	"gp/backend/db"
)

// AllFilters returns a slice of db.Artist containing any artist that matches any of the criteria.
func AllFilters(fv *Values, dataBase []*db.Artist) []*db.Artist {
	var results []*db.Artist = []*db.Artist{}

	for _, artist := range dataBase {
		if artist == nil {
			continue
		}
		if byCreation(fv.Creation, artist) &&
			byFirstAlbum(fv.FirstAlbum, artist) &&
			byMembers(fv.Members, artist) &&
			byLocations(fv.Locations, artist) {
			results = append(results, artist)
		}
	}
	return results
}
