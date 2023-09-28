// Package weather gives the current
// weather condition at the given location.
package weather

// CurrentCondition is the current weather condition at the given location.
var CurrentCondition string

// CurrentLocation is the location where the weather condition will be given.
var CurrentLocation string

// Forecast returns the current weather condition at the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
