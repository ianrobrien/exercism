package raindrops

import "strconv"

func Convert(number int) string {
	var hasFactor bool
	result := ""
	if number%3 == 0 {
		result += "Pling"
		hasFactor = true
	}
	if number%5 == 0 {
		result += "Plang"
		hasFactor = true
	}
	if number%7 == 0 {
		result += "Plong"
		hasFactor = true
	}
	if !hasFactor {
		return strconv.Itoa(number)
	}
	return result
}
