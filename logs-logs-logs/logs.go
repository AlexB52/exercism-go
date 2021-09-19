package logs

import (
	"fmt"
	"strings"
	"regexp"
)

func parseLog(line string) []string {
	re := regexp.MustCompile(`\[(.*)\]:(.*)`)
	return re.FindStringSubmatch(line)
}

// Message extracts the message from the provided log line.
func Message(line string) string {
	match := parseLog(line)
	return strings.TrimSpace(match[2])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return len([]rune(Message(line)))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	match := parseLog(line)
	return strings.ToLower(match[1])
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
