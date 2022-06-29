package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

// - `-n` Print the line numbers of each matching line.
// - `-l` Print only the names of files that contain at least one matching line.
// - `-i` Match line using a case-insensitive comparison.
// - `-v` Invert the program -- collect all lines that fail to match the pattern.
// - `-x` Only match entire lines, instead of lines that contain a match.

type Options struct {
	lineNumber      bool
	filename        bool
	caseInsensitive bool
	mistmatches     bool
	entireLine      bool
}

func buildOptions(flags []string) (result Options) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			result.lineNumber = true
		case "-l":
			result.filename = true
		case "-i":
			result.caseInsensitive = true
		case "-v":
			result.mistmatches = true
		case "-x":
			result.entireLine = true
		}
	}
	return result
}

func Search(pattern string, flags, files []string) (result []string) {
	options := buildOptions(flags)

	matches, mismatches := []string{}, []string{}
	matchedFiles := map[string]bool{}

	for _, filename := range files {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		lineNumber := 0

		if options.caseInsensitive {
			pattern = fmt.Sprintf("(?i)%s", pattern)
		}

		if options.entireLine {
			pattern = fmt.Sprintf("^%s$", pattern)
		}

		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()

			re := regexp.MustCompile(pattern)
			fmt.Println("match pattern?", re.Match([]byte(line)))

			if !re.Match([]byte(line)) {
				if options.lineNumber {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}

				if len(files) > 1 {
					line = fmt.Sprintf("%s:%s", filename, line)
				}

				mismatches = append(mismatches, line)
				continue
			}

			matchedFiles[filename] = true

			if options.lineNumber {
				line = fmt.Sprintf("%d:%s", lineNumber, line)
			}

			if len(files) > 1 {
				line = fmt.Sprintf("%s:%s", filename, line)
			}

			matches = append(matches, line)
		}
	}

	fmt.Println("pattern", pattern)
	fmt.Println("flags", flags)
	fmt.Printf("files: %q\n", files)
	fmt.Printf("matchedFiles: %v\n", matchedFiles)
	fmt.Println("matches:", matches)

	for _, flag := range flags {
		switch flag {
		case "-l":
			result = []string{}
			for k, _ := range matchedFiles {
				result = append(result, k)
			}
			sort.Strings(result)
			return result
		case "-v":
			return mismatches
		}
	}
	return matches
}
