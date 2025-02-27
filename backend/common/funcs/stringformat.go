package commonfuncs

import "strings"

// Transforms "washington_dc-usa" to "Washington Dc, USA"
func FormatLocation(input string) string {
	cleaned := strings.ReplaceAll(input, "_", " ")
	cleaned = strings.ReplaceAll(cleaned, "-", ", ")
	words := strings.Fields(cleaned)

	for i, word := range words {
		switch strings.ToLower(word) {
		case "usa":
			words[i] = "USA"
		case "uae":
			words[i] = "UAE"
		default:
			if len(word) > 0 {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
	}

	return strings.Join(words, " ")
}
