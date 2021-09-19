package logs

import (
	"fmt"
	"strings"
	"regexp"
)

type Log struct {
	message string
	level string
}

func parseLog(line string) Log {
	re := regexp.MustCompile(`\[(.*)\]:(.*)`)
	match := re.FindStringSubmatch(line)

	return Log{
		message: strings.TrimSpace(match[2]),
		level: strings.ToLower(match[1]),
	}
}

// Message extracts the message from the provided log line.
func Message(line string) string {
	return parseLog(line).message
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return len([]rune(Message(line)))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	return parseLog(line).level
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
