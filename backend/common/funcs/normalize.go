package commonfuncs

import (
	"regexp"
	"strings"
)

func Normalize(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.ToLower(re.ReplaceAllString(s, " "))
}
