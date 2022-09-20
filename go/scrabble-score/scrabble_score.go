package scrabble

import "strings"

func Score(word string) int {
	score := 0
	upperWord := strings.ToUpper(word)
	for i := 0; i < len(word); i++ {
		score += letterValue(string(upperWord[i]))
	}
	return score
}

func letterValue(letter string) int {
	switch {
	case strings.Contains("AEIOULNRST", letter):
		return 1
	case strings.Contains("DG", letter):
		return 2
	case strings.Contains("BCMP", letter):
		return 3
	case strings.Contains("FHVWY", letter):
		return 4
	case strings.Contains("K", letter):
		return 5
	case strings.Contains("JX", letter):
		return 8
	case strings.Contains("QZ", letter):
		return 10
	}
	panic("Unknown letter")
}
