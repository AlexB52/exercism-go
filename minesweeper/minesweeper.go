package minesweeper

// Annotate returns an annotated board
type Incrementer interface {
	increment()
}

type Mine struct {
	x, y  int
	Value interface{}
}

type Position struct {
	x, y  int
	Value interface{}
}

func (p *Position) increment() {
	p.Value++
}

func (p *Mine) increment() {}

func Annotate(board []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	matrix := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		matrix[i] = make([]int, len(board[0]))
	}

	for i, row := range board {
		for j, col := range row {
			if col != '*' {
				continue
			}
			matrix[i][j]++
		}
	}

	return board
}
