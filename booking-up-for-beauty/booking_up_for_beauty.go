package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	time, _ := time.Parse("1/2/2006 15:04:05", date)
	return time
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	timeReference, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(timeReference)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	timeReference, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return 12 <= timeReference.Hour() && timeReference.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	timeReference, _ := time.Parse("1/2/2006 15:04:05", date)
	timeFormatted := timeReference.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", timeFormatted)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
