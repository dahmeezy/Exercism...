// Package weather provides tools to check the weather and forcast it.
package weather

var (
// CurrentCondition repesents the current weather condition.
	CurrentCondition string
// CurrentLocation represents the city.
	CurrentLocation  string
)
// Forecast returns string values that states the current weather condition of the city passed.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
