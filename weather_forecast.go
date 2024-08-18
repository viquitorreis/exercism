// Package weather provides tools to get
// current condition and current location.
package weather

// CurrentCondition represents the condition from cities.
var CurrentCondition string

// CurrentLocation represents the location of a city.
var CurrentLocation string

// Forecast returns the current location and current condition from cities.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
