package search

import (
	commonfuncs "gp/backend/common/funcs"
	"strconv"
	"strings"
)

func compareStrings(query, name string) bool {
	return strings.Contains(commonfuncs.Normalize(name), query) ||
		strings.Contains(query, commonfuncs.Normalize(name))
}

func checkMembers(query string, members []string) bool {
	for _, member := range members {
		if compareStrings(member, query) {
			return true
		}
	}
	return false
}

func checkCreationDate(query string, dateStr int) bool {
	dateInt := strconv.Itoa(dateStr)
	if compareStrings(dateInt, query) {
		return true
	}
	return false
}

// Range over artist locations and if
func CheckRelations(query string, locations map[string][]string) bool {
	for location := range locations {
		if compareStrings(location, query) {
			return true
		}
	}
	return false
}
