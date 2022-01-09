package matrix

// Define the Matrix type here.

import (
	"fmt"
	"strings"
	"strconv"
)

type Matrix struct {}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	numbers := make([][]int, len(rows))

	for i, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		cells := make([]int, len(cols))

		for j, cell := range cols {
			integer, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			cells[j] = integer
		}
		numbers[i] = cells
	}

	fmt.Println(numbers)
	return &Matrix{}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m *Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
