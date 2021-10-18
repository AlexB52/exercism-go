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
	clock.hour, clock.minute = ConvertMinutes(clock.TotalMinutes() + minutes)
	return clock
}

func (clock Clock) Subtract(minutes int) Clock {
	clock.hour, clock.minute = ConvertMinutes(clock.TotalMinutes() - minutes)
	return clock
}

func (clock Clock) TotalMinutes() int {
	return clock.hour*60 + clock.minute
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
