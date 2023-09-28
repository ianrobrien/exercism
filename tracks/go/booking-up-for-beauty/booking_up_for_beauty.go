package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("error")
	}
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedTime, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic("error")
	}
	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic("error")
	}
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic("error")
	}
	return fmt.Sprintf(
		"You have an appointment on %s at %s",
		parsedTime.Format("Monday, January 2, 2006,"),
		parsedTime.Format("15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedTime, err := time.Parse("1/2/2006 15:04:05", "9/15/2020 00:00:00")
	if err != nil {
		panic("error")
	}
	return time.Date(time.Now().Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, time.UTC)
}
