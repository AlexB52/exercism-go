package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	start, end := DayRange(week, year, month)
	for day := start; day <= end; day++ {
		if weekday == time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Weekday() {
			return day
		}
	}
	return 0
}

func DayRange(week WeekSchedule, year int, month time.Month) (start int, end int) {
	switch week {
	case First:
		return 1, 8
	case Second:
		return 8, 15
	case Third:
		return 15, 22
	case Fourth:
		return 22, 29
	case Teenth:
		return 13, 19
	case Last:
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
		return lastDay - 6, lastDay
	default:
		return 0, 0
	}
}
