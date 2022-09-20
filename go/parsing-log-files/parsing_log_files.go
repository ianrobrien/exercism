package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC]|^\[DBG]|^\[INF]|^\[WRN]|^\[ERR]|^\[FTL]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[!~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*password.*"`)

	count := 0
	for i := range lines {
		password := re.FindString(strings.ToLower(lines[i]))
		if password != "" {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return ""
}

func TagWithUserName(lines []string) []string {
	return []string{}

}
