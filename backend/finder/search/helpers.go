package search

import (
	"strings"
)

func isContained(query, name string) bool {
	return strings.Contains(name, query) ||
		strings.Contains(query, name)
}

func startsWith(query, name string) bool {
	return strings.HasPrefix(name, query) || name == query
}
