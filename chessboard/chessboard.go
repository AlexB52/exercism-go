package chessboard

// import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	for _, position := range cb[rank] {
		if (position) { ret += 1 }
	}

	return
}

// // CountInFile returns how many squares are occupied in the chessboard,
// // within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	if file < 1 || file > len(cb) {
		return 0
	}

	for _, rank := range cb {
		if (rank[file-1]) { ret ++ }
	}

	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	for _, rank := range cb {
		ret += len(rank)
	}

	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for index, _ := range cb {
		ret += CountInRank(cb, index)
	}

	return
}
