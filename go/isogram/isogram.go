package isogram

import "strings"

func IsIsogram(word string) bool {
	letterMap := make(map[string]bool)

	normalizedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	for i := range normalizedWord {
		key := string(normalizedWord[i])
		if key == "-" {
			continue
		}
		value, exists := letterMap[key]
		if exists {
			return false
		}
		letterMap[key] = value
	}

	return true
}
