package series

func All(n int, s string) (result []string) {
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
