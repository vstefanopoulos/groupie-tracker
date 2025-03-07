package search

import (
	"gp/backend/db"
	"gp/backend/finder/shared"
	"strconv"
)

func lookInNames(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	name := shared.Normalize(artist.Name)
	if shared.StartsWith(query, name) {
		addRank = 100
		match = true
	} else if shared.IsContained(query, name) {
		addRank = 10
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.Name = true
		} else {
			result = &SearchItem{
				Artist: artist,
				Rank:   addRank,
				Name:   true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInAlbum(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	firstAlbum := shared.Normalize(artist.FirstAlbum)
	if shared.StartsWith(query, firstAlbum) {
		addRank = 90
		match = true

	} else if shared.IsContained(query, firstAlbum) {
		addRank = 10
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.Album = true
		} else {
			result = &SearchItem{
				Artist: artist,
				Rank:   addRank,
				Album:  true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInMembers(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	for _, member := range artist.Members {
		match := false
		member = shared.Normalize(member)
		if shared.StartsWith(query, member) {
			addRank = 80
			match = true
		} else if shared.IsContained(query, member) {
			addRank = 10
			match = true
		}

		if match {
			if result, exists := results[artist.ID]; exists {
				result.Rank += addRank
				result.Members = true
			} else {
				result = &SearchItem{
					Artist:  artist,
					Rank:    addRank,
					Members: true,
				}
				results[artist.ID] = result
			}
		}
	}
}

func lookInCreation(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	var match bool
	dateStr := shared.Normalize(strconv.Itoa(artist.CreationDate))

	if shared.StartsWith(query, dateStr) {
		addRank = 70
		match = true
	} else if shared.IsContained(query, dateStr) {
		addRank = 10
		match = true
	}

	if match {
		if result, exists := results[artist.ID]; exists {
			result.Rank += addRank
			result.Creation = true
		} else {
			result = &SearchItem{
				Artist:   artist,
				Rank:     addRank,
				Creation: true,
			}
			results[artist.ID] = result
		}
	}
}

func lookInLocations(query string, artist *db.Artist, results map[int]*SearchItem) {
	var addRank int
	for location := range artist.Relation.DatesLocations {
		match := false
		location = shared.Normalize(location)
		if shared.StartsWith(query, location) {
			addRank = 60
			match = true
		} else if shared.IsContained(query, location) {
			addRank = 10
			match = true
		}

		if match {
			if result, exists := results[artist.ID]; exists {
				result.Rank += addRank
				result.Locations = true
			} else {
				result = &SearchItem{
					Artist:    artist,
					Rank:      addRank,
					Locations: true,
				}
				results[artist.ID] = result
			}
		}
	}
}
