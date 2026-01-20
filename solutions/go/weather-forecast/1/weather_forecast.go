// Package weather provides functionalities to get weather forecasts.
package weather

var (
	// CurrentCondition holds the current weather condition.
	CurrentCondition string
	// CurrentLocation holds the current location for the weather forecast.
	CurrentLocation string
)

// Forecast returns a formatted string with the current location and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
