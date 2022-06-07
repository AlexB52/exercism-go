package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\](\s\w)+`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[^a-z]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*[pP][aA][sS][sS][wW][oO][rR][dD].*"`)
	return len(re.FindAllString(strings.Join(lines, "\n"), -1))
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (result []string) {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		if matches != nil {
			l = fmt.Sprintf("[USR] %s %s", matches[1], l)
		}
		result = append(result, l)
	}
	return result
}
