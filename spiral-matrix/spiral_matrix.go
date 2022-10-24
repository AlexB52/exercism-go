package spiralmatrix

type Position struct {
	x, y int
}

type Direction struct {
	move func(Position) Position
	next *Direction
}

func SpiralMatrix(size int) [][]int {
	var result = make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}

	position := Position{0, 0}
	direction := BuildDirection()

	for i := 1; i <= size*size; i++ {
		result[position.x][position.y] = i

		for turns := 0; turns <= 1; turns++ {
			newPosition := direction.move(position)
			if ValidMove(result, newPosition) {
				position = newPosition
				break
			} else {
				direction = direction.next
			}
		}
	}

	return result
}

func ValidMove(grid [][]int, p Position) bool {
	return 0 <= p.x && p.x < len(grid) && 0 <= p.y && p.y < len(grid) && grid[p.x][p.y] == 0
}

func BuildDirection() *Direction {
	up := &Direction{
		move: func(p Position) Position { return Position{p.x - 1, p.y} },
	}

	left := &Direction{
		move: func(p Position) Position { return Position{p.x, p.y - 1} },
		next: up,
	}

	down := &Direction{
		move: func(p Position) Position { return Position{p.x + 1, p.y} },
		next: left,
	}

	right := &Direction{
		move: func(p Position) Position { return Position{p.x, p.y + 1} },
		next: down,
	}

	up.next = right
	return right
}
