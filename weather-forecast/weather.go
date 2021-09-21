// Package: provides a function to display the weather forecast
package weather

var (
	// this is the current condition of the User
	CurrentCondition string

	// this is the current location of the User
	CurrentLocation string
)

// Forecast formats the forecast message
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
