package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	hour, minute = ConvertMinutes(hour*60 + minute)
	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute + minutes)
}

func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.hour, clock.minute - minutes)
}

func ConvertMinutes(minutes int) (hour, minute int) {
	hour = minutes / 60 % 24
	minute = minutes % 60

	if minute < 0 {
		minute += 60
		hour--
	}

	if hour < 0 {
		hour += 24
	}

	return hour, minute
}
