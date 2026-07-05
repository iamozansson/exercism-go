// Package weather provides forecast information
// about cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition of the user's city.
	CurrentCondition string
	// CurrentLocation represents the current location of the user.
	CurrentLocation string
)

// Forecast returns a formatted string describing the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
