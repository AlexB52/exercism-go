package diffsquares

func SquareOfSum(n int) (result int) {
	for i := 0; i <= n; i++ {
		result += i
	}

	return result * result
}

func SumOfSquares(n int) (result int) {
	for i := 0; i <= n; i++ {
		result += i * i
	}

	return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
