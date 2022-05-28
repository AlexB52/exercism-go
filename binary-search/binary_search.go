package binarysearch

func SearchInts(list []int, key int) int {
	var start, end = 0, len(list) - 1
	for start <= end {
		index := (end + start) / 2

		if list[index] > key {
			end = index - 1
			continue
		}

		if list[index] < key {
			start = index + 1
			continue
		}

		return index
	}

	return -1
}
