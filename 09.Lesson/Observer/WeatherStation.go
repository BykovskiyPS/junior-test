// Observe pattern
// Iot  system. Weather station

package Observer

import (
	"fmt"
)

type Publisher interface {
	Attach(observer Observer)
	Remove(observer Observer)
	Notify()
}

type Observer interface {
	Update(temp float64, pressure float64)
	Display()
}

type WeatherStation struct {
	observers []Observer
	temp      float64
	pressure  float64
}

func NewPublisher() Publisher {
	return &WeatherStation{}
}

func (s *WeatherStation) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *WeatherStation) Remove(observer Observer) {
	for i, v := range s.observers {
		if v == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *WeatherStation) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.temp, s.pressure)
	}
}

func (s *WeatherStation) setMeasurements(temp float64, pressure float64) {
	s.temp = temp
	s.pressure = pressure
	s.Notify()
}

type ConcreteObserver struct {
	temp           float64
	pressure       float64
	WeatherStation Publisher
}

func (c *ConcreteObserver) Update(temp float64, pressure float64) {
	c.temp = temp
	c.pressure = pressure
	c.Display()
}

func (c *ConcreteObserver) Display() {
	fmt.Println("Temperature: ", c.temp, "Pressure: ", c.pressure)
}
