package shared

import (
	"regexp"
	"strings"
)

func IsContained(query, name string) bool {
	return strings.Contains(name, query) || strings.Contains(query, name)
}

func StartsWith(query, name string) bool {
	return strings.HasPrefix(name, query) || name == query
}

func Normalize(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.ToLower(re.ReplaceAllString(s, " "))
}
