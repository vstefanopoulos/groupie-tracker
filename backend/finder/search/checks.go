package search

import (
	commonfuncs "gp/backend/common/funcs"
	"gp/backend/db"
	"strconv"
)

// Range over artist locations and if
func CheckRelations(query string, locations map[string][]string) bool {
	for location := range locations {
		if isContained(query, commonfuncs.Normalize(location)) {
			return true
		}
	}
	return false
}

func lookInNames(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	name := commonfuncs.Normalize(artist.Name)
	if startsWith(query, name) {
		addRank = 100
		match = true
	} else if isContained(query, name) {
		addRank = 50
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.artist = true
		} else {
			result = &SearchItem{
				Artist: artist,
				Rank:   addRank,
				artist: true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInAlbum(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	firstAlbum := commonfuncs.Normalize(artist.FirstAlbum)
	if startsWith(query, firstAlbum) {
		addRank = 90
		match = true

	} else if isContained(query, firstAlbum) {
		addRank = 40
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.album = true
		} else {
			result = &SearchItem{
				Artist: artist,
				Rank:   addRank,
				album:  true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInMembers(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	for _, member := range artist.Members {
		match := false
		member = commonfuncs.Normalize(member)
		if startsWith(query, member) {
			addRank = 80
			match = true
		} else if isContained(query, member) {
			addRank = 30
			match = true
		}

		if match {
			if result, exists := results[artist.ID]; exists {
				result.Rank += addRank
				result.members = true
			} else {
				result = &SearchItem{
					Artist:  artist,
					Rank:    addRank,
					members: true,
				}
				results[artist.ID] = result
			}
		}
	}
}

func lookInCreation(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	dateStr := commonfuncs.Normalize(strconv.Itoa(artist.CreationDate))

	if startsWith(query, dateStr) {
		addRank = 70
		match = true
	} else if isContained(query, dateStr) {
		addRank = 20
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.creation = true
		} else {
			result = &SearchItem{
				Artist:   artist,
				Rank:     addRank,
				creation: true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInLocations(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	for location := range artist.Relation.DatesLocations {
		match := false
		location = commonfuncs.Normalize(location)
		if startsWith(query, location) {
			addRank = 60
			match = true
		} else if isContained(query, location) {
			addRank = 10
			match = true
		}

		if match {
			if result, exists := results[artist.ID]; exists {
				result.Rank += addRank
				result.locations = true
			} else {
				result = &SearchItem{
					Artist:    artist,
					Rank:      addRank,
					locations: true,
				}
				results[artist.ID] = result
			}
		}
	}
}
