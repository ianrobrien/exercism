package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	hour := h
	minute := m
	for minute >= 60 {
		minute -= 60
		hour++
	}
	for hour >= 24 {
		hour -= 24
	}
	return Clock{
		hour:   hour,
		minute: minute,
	}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	return c
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
