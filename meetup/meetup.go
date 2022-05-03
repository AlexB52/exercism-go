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
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)

	for day := 1; day <= t.Day(); day++ {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		switch week {
		case First:
			if weekday == date.Weekday() && day > 0 {
				return day
			}
		case Second:
			if weekday == date.Weekday() && day > 7 {
				return day
			}
		case Third:
			if weekday == date.Weekday() && day > 14 {
				return day
			}
		case Fourth:
			if weekday == date.Weekday() && day > 21 {
				return day
			}
		case Teenth:
			if weekday == date.Weekday() && day >= 13 && day <= 19 {
				return day
			}
		case Last:
			if weekday == date.Weekday() && day > t.Day()-7 {
				return day
			}
		}
	}
	return 0
}
