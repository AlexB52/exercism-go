package wordsearch

import (
	"fmt"
	"regexp"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := map[string][2][2]int{}

	var matrix = make([][]rune, len(puzzle))
	for i := 0; i < len(puzzle); i++ {
		matrix[i] = []rune(puzzle[i])
	}

	for _, word := range words {
		for i, line := range puzzle {
			if match := regexp.MustCompile(word).FindStringIndex(line); match != nil {
				result[word] = LeftToRight(match, i)
			}

			if match := regexp.MustCompile(Reverse(word)).FindStringIndex(line); match != nil {
				result[word] = RightToLeft(match, i)
			}
		}

		for i, line := range TransposePuzzle(puzzle) {
			if match := regexp.MustCompile(word).FindStringIndex(line); match != nil {
				result[word] = LeftToRightReverse(match, i)
			}

			if match := regexp.MustCompile(Reverse(word)).FindStringIndex(line); match != nil {
				result[word] = RightToLeftReverse(match, i)
			}
		}

		fmt.Println("word", word)
		for i := 0; i < len(matrix); i++ {
			if match := regexp.MustCompile(word).FindStringIndex(TopBottomDiagonalLine(i, 0, matrix)); match != nil {
				result[word] = LeftToRightDiagonal(match, i)
			}

			if match := regexp.MustCompile(Reverse(word)).FindStringIndex(TopBottomDiagonalLine(i, 0, matrix)); match != nil {
				result[word] = RightToLeftDiagonal(match, i)
			}
		}

		for i := 0; i < len(matrix); i++ {
			if match := regexp.MustCompile(word).FindStringIndex(BottomTopDiagonalLine(i, 0, matrix)); match != nil {
				result[word] = BottomLeftToTopRight(match, i)
			}
		}

		for j := 0; j < len(matrix[0]); j++ {
			if match := regexp.MustCompile(Reverse(word)).FindStringIndex(BottomTopDiagonalLine(len(matrix)-1, j, matrix)); match != nil {
				result[word] = TopRightBottomLeftReverse(match, j)
			}
		}

		fmt.Println()
		// TopBottomDiagonalLine(0, 0, matrix)
	}

	return result, nil
}

func TopBottomDiagonalLine(x, y int, matrix [][]rune) string {
	var result []rune
	for x < len(matrix) && y < len(matrix[0]) {
		result = append(result, matrix[x][y])
		x++
		y++
	}
	fmt.Println("result", string(result))
	return string(result)
}

func BottomTopDiagonalLine(x, y int, matrix [][]rune) string {
	var result []rune
	for x >= 0 && y < len(matrix[0]) {
		result = append(result, matrix[x][y])
		x--
		y++
	}
	return string(result)
}

func LeftToRight(match []int, i int) (result [2][2]int) {
	result[0][0] = match[0]
	result[0][1] = i
	result[1][0] = match[1] - 1
	result[1][1] = i
	return result
}

func RightToLeft(match []int, i int) (result [2][2]int) {
	result[0][0] = match[1] - 1
	result[0][1] = i
	result[1][0] = match[0]
	result[1][1] = i
	return result
}

func LeftToRightReverse(match []int, i int) (result [2][2]int) {
	result[0][0] = i
	result[0][1] = match[0]
	result[1][0] = match[1] - 1
	result[1][1] = i
	return result
}

func RightToLeftReverse(match []int, i int) (result [2][2]int) {
	result[0][0] = i
	result[0][1] = match[1] - 1
	result[1][0] = i
	result[1][1] = match[0]
	return result
}

func LeftToRightDiagonal(match []int, i int) (result [2][2]int) {
	result[0][0] = match[0]
	result[0][1] = i
	result[1][0] = match[1] - 1
	result[1][1] = match[1] - 1
	return result
}

func BottomLeftToTopRight(match []int, i int) (result [2][2]int) {
	result[0][0] = match[0]
	result[0][1] = match[1] - 1
	result[1][0] = match[1] - 1
	result[1][1] = match[0]
	return result
}

func TopRightToBottomLeft(match []int, i int) (result [2][2]int) {
	result[0][0] = match[0]
	result[0][1] = match[1] - 1
	result[1][0] = match[1] - 1
	result[1][1] = match[0]
	return result
}

func RightToLeftDiagonal(match []int, i int) (result [2][2]int) {
	result[0][0] = match[1] - 1
	result[0][1] = i + match[1] - 1
	result[1][0] = match[0]
	result[1][1] = i + match[0]
	return result
}

func TopRightBottomLeftReverse(match []int, i int) (result [2][2]int) {
	result[0][0] = i + match[1] - 1
	result[0][1] = match[1]
	result[1][0] = match[1] - 1
	result[1][1] = i + match[1]
	return result
}

func Reverse(word string) (result string) {
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}

func TransposePuzzle(puzzle []string) []string {
	var result = make([][]rune, len(puzzle[0]))
	for i := 0; i < len(puzzle[0]); i++ {
		result[i] = make([]rune, len(puzzle))
	}

	for i, line := range puzzle {
		for j, char := range []rune(line) {
			result[j][i] = char
		}
	}

	var reverse = make([]string, len(puzzle[0]))
	for i, line := range result {
		reverse[i] = string(line)
	}

	return reverse
}
