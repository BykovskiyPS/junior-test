package Observer

import "testing"

func TestWeatherStation(t *testing.T) {
	var weather WeatherStation
	var conditionDisplay ConcreteObserver

	weather.Attach(&conditionDisplay)

	weather.setMeasurements(4.4, 5.5)
	weather.setMeasurements(6.6, 9.2)
	weather.setMeasurements(7.7, 7.7)
}
