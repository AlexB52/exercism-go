package binarysearch

func SearchInts(list []int, key int) int {
	var index int
	var start, end = 0, len(list)
	for len(list[start:end]) > 1 {
		index = (end + start) / 2
		if list[index] > key {
			end = index
		} else if list[index] < key {
			start = index
		} else if list[index] == key {
			return index
		}
	}

	if len(list[start:end]) == 1 && list[start:end][0] == key {
		return start
	}

	return -1
}
