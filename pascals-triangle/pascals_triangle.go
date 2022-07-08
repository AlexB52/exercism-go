package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)
	result[0] = []int{1}
	for i := 1; i < n; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < len(result[i-1]); j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
