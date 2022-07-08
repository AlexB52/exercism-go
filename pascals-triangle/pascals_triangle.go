package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)
	result[0] = []int{1}

	for i := 1; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < len(result[i-1]); j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = row
	}
	return result
}
