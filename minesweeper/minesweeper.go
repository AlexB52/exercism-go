package minesweeper

import (
	"fmt"
)

// Annotate returns an annotated board

func Annotate(board []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	mineBoard := make([][]rune, len(board))
	for i := 0; i < len(board); i++ {
		mineBoard[i] = make([]rune, len(board[0]))
	}

	for i, row := range board {
		for j, col := range row {
			if col != '*' {
				continue
			}
			mineBoard[i][j] = '*'
			Increment(mineBoard, i-1, j-1)
			Increment(mineBoard, i-1, j)
			Increment(mineBoard, i-1, j+1)
			Increment(mineBoard, i, j-1)
			Increment(mineBoard, i, j+1)
			Increment(mineBoard, i+1, j-1)
			Increment(mineBoard, i+1, j)
			Increment(mineBoard, i+1, j+1)
		}
	}

	return PrintResult(mineBoard)
}

func Increment(board [][]rune, i, j int) {
	x, y := len(board)-1, len(board[0])-1
	if i > x || i < 0 || j > y || j < 0 || board[i][j] == '*' {
		return
	}
	board[i][j]++
}

func PrintResult(board [][]rune) []string {
	var result = make([]string, len(board))
	for i, line := range board {
		for _, char := range line {
			var value string
			switch char {
			case '*':
				value = "*"
			case 0:
				value = " "
			default:
				value = fmt.Sprint(char - 0)
			}
			result[i] += value
		}
	}
	return result
}
