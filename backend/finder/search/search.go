package search

import (
	commonfuncs "gp/backend/common/funcs"
	"gp/backend/db"
	"strings"
)

// Parses query first word and conducts search according to the flag:
//   - "artist-band " by name
//   - "member " by group members
//   - "first-album " by first album release date
//   - "creation-date " by creation date
//   - "location " by tour location
func Search(query string, dataBase []*db.Artist) []*db.Artist {
	query = commonfuncs.Normalize(query)
	searchType := strings.Fields(query)
	if len(searchType) == 1 {
		return searchAll(query, dataBase)
	}

	fmtQuery := strings.Join(searchType[1:], " ")

	switch searchType[0] {
	case "artist":
		return searchName(fmtQuery, dataBase)
	case "members":
		return searchMember(fmtQuery, dataBase)
	case "album":
		return searchAlbum(fmtQuery, dataBase)
	case "creation":
		return searchCreation(fmtQuery, dataBase)
	case "locations":
		return searchLocation(fmtQuery, dataBase)
	default:
		return searchAll(query, dataBase)
	}
}

func searchAll(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {
		if compareStrings(artist.Name, query) ||
			compareStrings(artist.FirstAlbum, query) ||
			checkMembers(query, artist.Members) ||
			checkCreationDate(query, artist.CreationDate) ||
			CheckRelations(query, artist.Relation.DatesLocations) {
			feed = append(feed, artist)
		}
	}
	return feed
}

func searchName(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {

		if compareStrings(query, commonfuncs.Normalize(artist.Name)) {
			feed = append(feed, artist)
		}
	}
	return feed
}

func searchLocation(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {
		if CheckRelations(query, artist.Relation.DatesLocations) {
			feed = append(feed, artist)
		}
	}
	return feed
}

func searchCreation(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {
		if checkCreationDate(query, artist.CreationDate) {
			feed = append(feed, artist)
		}
	}
	return feed
}

func searchAlbum(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {
		if compareStrings(query, artist.FirstAlbum) {
			feed = append(feed, artist)
		}
	}
	return feed
}

func searchMember(query string, dataBase []*db.Artist) []*db.Artist {
	feed := []*db.Artist{}
	for _, artist := range dataBase {
		if checkMembers(query, artist.Members) {
			feed = append(feed, artist)
		}
	}
	return feed
}
