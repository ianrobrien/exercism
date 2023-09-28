package parsinglogfiles

import (
	"fmt"
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
	re := regexp.MustCompile(`end-of-line[0-9]*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)

	var out []string

	for i := range lines {
		var line string
		tokens := re.FindStringSubmatch(lines[i])
		if len(tokens) == 2 {
			line = fmt.Sprintf("[USR] %s %s", tokens[1], lines[i])
		} else {
			line = lines[i]
		}
		out = append(out, line)
	}

	return out

}
