package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	// Spaces are allowed in the input, but they should be stripped before checking
	normalizedId := strings.ReplaceAll(id, " ", "")

	// Strings of length 1 or less are not valid.
	if len(normalizedId) <= 1 {
		return false
	}

	//	All other non-digit characters are disallowed.
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(normalizedId) {
		return false
	}

	var result string
	for i := len(normalizedId) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(normalizedId[i]))
		if (len(normalizedId)-i)%2 == 0 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		result = strconv.Itoa(digit) + result
	}

	var sum int
	for i := 0; i < len(result); i++ {
		value, _ := strconv.Atoi(string(result[i]))
		sum += value
	}

	return sum%10 == 0
}
