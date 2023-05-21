// Package weather forecasts current weather conditions for various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition of the location set.
var CurrentCondition string
// CurrentLocation stores the current location for which the weather condition is being forecasted.
var CurrentLocation string

// Forecast returns a string containing the current weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
