package clock

import "fmt"

// Clock represents a standard 24-hour clock with values from 00:00 to 23:59
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	minute := m
	hour := h

	for minute < 0 {
		minute += 60
		hour -= 1
	}
	for minute >= 60 {
		minute -= 60
		hour += 1
	}

	for hour < 0 {
		hour += 24
	}
	for hour >= 24 {
		hour -= 24
	}

	return Clock{hour: hour, minute: minute}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
