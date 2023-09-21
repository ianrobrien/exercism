package chessboard

import (
	"strings"
)

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	mappedFile, exists := cb[strings.ToUpper(file)]
	if !exists {
		return 0
	}

	occupiedCount := 0
	for _, occupied := range mappedFile {
		if occupied {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	occupiedCount := 0
	for _, letter := range "ABCDEFGH" {
		parsedLetter := string(letter)
		file, _ := cb[parsedLetter]
		occupied := file[rank-1]
		if occupied {
			occupiedCount++
		}
	}
	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}
	return count
}
