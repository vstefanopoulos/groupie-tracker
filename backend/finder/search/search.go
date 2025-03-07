package search

import (
	"gp/backend/db"
	"gp/backend/finder/shared"
	"sort"
	"strings"
)

type SearchItem struct {
	Artist                                    *db.Artist
	Rank                                      int
	Name, Members, Album, Creation, Locations bool
}

type SearchTags struct {
	query                                       string
	artist, members, album, creation, locations bool
}

// Parses query first word and conducts search according to the flag:
//   - "artist-band " by name
//   - "member " by group members
//   - "first-album " by first album release date
//   - "creation-date " by creation date
//   - "location " by tour location
func Search(query string, dataBase []*db.Artist) []*SearchItem {
	query = shared.Normalize(query)

	allTags := &SearchTags{
		query:     query,
		artist:    true,
		members:   true,
		album:     true,
		creation:  true,
		locations: true,
	}

	searchType := strings.Fields(query)
	if len(searchType) == 1 {
		return sortByRank(searchByTag(allTags, dataBase))
	}

	tags := &SearchTags{query: strings.Join(searchType[1:], " ")}

	switch searchType[0] {
	case "artist":
		tags.artist = true
	case "members":
		tags.members = true
	case "album":
		tags.album = true
	case "creation":
		tags.creation = true
	case "locations":
		tags.locations = true
	default:
		tags = allTags
	}

	results := searchByTag(tags, dataBase)
	ranked := sortByRank(results)
	return ranked
}

func searchByTag(query *SearchTags, dataBase []*db.Artist) map[int]*SearchItem {
	results := make(map[int]*SearchItem)

	for _, artist := range dataBase {
		if query.artist {
			lookInNames(query.query, artist, results)
		}

		if query.album {
			lookInAlbum(query.query, artist, results)
		}

		if query.members {
			lookInMembers(query.query, artist, results)
		}

		if query.creation {
			lookInCreation(query.query, artist, results)
		}

		if query.locations {
			lookInLocations(query.query, artist, results)
		}
	}

	return results
}

func sortByRank(results map[int]*SearchItem) []*SearchItem {
	feed := []*SearchItem{}

	if results == nil {
		return feed
	}

	for _, result := range results {
		feed = append(feed, result)
	}

	sort.Slice(feed, func(i, j int) bool {
		return feed[i].Rank > feed[j].Rank
	})

	return feed
}
