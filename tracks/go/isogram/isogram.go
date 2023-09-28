package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letterMap := make(map[rune]struct{})

	runes := []rune(strings.ToLower(word))

	for i := range runes {
		key := runes[i]
		if !unicode.IsLetter(key) {
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
