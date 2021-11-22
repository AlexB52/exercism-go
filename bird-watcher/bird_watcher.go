package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for _, count := range birdsPerDay {
		result += count
	}
	return result
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	offset := 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[offset : offset+7])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
