package isogram

import "strings"

func IsIsogram(word string) bool {
	letterMap := make(map[string]struct{})

	normalizedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	for i := range normalizedWord {
		key := string(normalizedWord[i])
		if key == "-" {
			continue
		}
		_, exists := letterMap[key]
		if exists {
			return false
		}
		letterMap[key] = struct{}{}
	}

	return true
}
