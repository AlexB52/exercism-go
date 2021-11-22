package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for _, count := range birdsPerDay {
		result += count
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var result int
	startingIndex := 7*(week-1)
	for i := startingIndex; i < startingIndex+7; i++ {
		result += birdsPerDay[i]
	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
