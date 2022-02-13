package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.

type Matrix struct {
	columnNumber int
	cells        []int
}

func New(s string) (*Matrix, error) {
	// var result []int
	rows := strings.Split(s, "\n")
	var columnNumber int
	var numbers []int

	for _, cell := range rows {
		cells := strings.Split(strings.TrimSpace(cell), " ")

		if columnNumber != 0 && columnNumber != len(cells) {
			return nil, errors.New("Invalid matrix columns")
		}

		columnNumber = len(cells)

		for _, strNumber := range cells {
			number, err := strconv.Atoi(strNumber)
			if err != nil {
				return nil, errors.New("Invalid input passed")
			}

			numbers = append(numbers, number)
		}
	}

	return &Matrix{columnNumber, numbers}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	numberOfRows := len(m.cells) / m.columnNumber
	result := make([][]int, m.columnNumber)

	for i := 0; i < m.columnNumber; i++ {
		result[i] = make([]int, numberOfRows)
	}

	for index, cell := range m.cells {
		i, j := index/m.columnNumber, index%m.columnNumber
		result[j][i] = cell
	}

	return result
}

func (m *Matrix) Rows() [][]int {
	numberOfRows := len(m.cells) / m.columnNumber
	result := make([][]int, numberOfRows)

	for i := 0; i < numberOfRows; i++ {
		result[i] = make([]int, m.columnNumber)
	}

	for index, cell := range m.cells {
		i, j := index/m.columnNumber, index%m.columnNumber
		result[i][j] = cell
	}

	return result
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
