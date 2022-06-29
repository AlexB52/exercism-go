package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Options struct {
	lineNumber      bool
	filename        bool
	caseInsensitive bool
	mismatches      bool
	entireLine      bool
	multipleFile    bool
	regex           *regexp.Regexp
}

type Line struct {
	Filename string
	Content  string
	Number   int
}

func Search(pattern string, flags, files []string) []string {
	options := buildOptions(pattern, flags, files)
	result, matched := []string{}, map[string]bool{}

	for _, filename := range files {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		lineNumber := 0

		for scanner.Scan() {
			lineNumber++
			line := Line{Filename: filename, Content: scanner.Text(), Number: lineNumber}

			if options.Match(line) {
				formattedLine := options.Format(line)
				if _, ok := matched[formattedLine]; !ok {
					matched[formattedLine] = true
					result = append(result, formattedLine)
				}
			}
		}
	}
	return result
}

func buildOptions(pattern string, flags, files []string) (result Options) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			result.lineNumber = true
		case "-l":
			result.filename = true
		case "-i":
			result.caseInsensitive = true
		case "-v":
			result.mismatches = true
		case "-x":
			result.entireLine = true
		}
	}
	result.multipleFile = len(files) > 1
	result.regex = regexp.MustCompile(result.BuildPattern(pattern))
	return result
}

func (o Options) Match(l Line) (result bool) {
	result = o.regex.Match([]byte(l.Content))
	if o.mismatches {
		result = !result
	}
	return result
}

func (o Options) Format(l Line) (result string) {
	if o.filename {
		return l.Filename
	}

	result = l.Content
	if o.lineNumber {
		result = fmt.Sprintf("%d:%s", l.Number, result)
	}
	if o.multipleFile {
		result = fmt.Sprintf("%s:%s", l.Filename, result)
	}
	return result
}

func (o Options) BuildPattern(pattern string) string {
	if o.caseInsensitive {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}
	if o.entireLine {
		pattern = fmt.Sprintf("^%s$", pattern)
	}
	return pattern
}
