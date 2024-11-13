package application

import (
	"context"
	"errors"
)

type Weather struct {
	celsius    float64
	fahrenheit float64
	kelvin     float64
}

type IWeatherRepository interface {
	GetTemperature(ctx context.Context, coordinate *Coordinate) (*Weather, error)
}

func NewWeather(tempCelsius float64) (*Weather, error) {
	w := &Weather{
		celsius:    tempCelsius,
		kelvin:     tempCelsius + 273.15,
		fahrenheit: tempCelsius*1.8 + 32,
	}

	if e := w.validate(); e != nil {
		return nil, e
	}

	return w, nil
}

func (w *Weather) Celsius() float64 {
	return w.celsius
}

func (w *Weather) Fahrenheit() float64 {
	return w.fahrenheit
}

func (w *Weather) Kelvin() float64 {
	return w.kelvin
}

func (w *Weather) validate() error {
	if w.celsius < -273.15 {
		return errors.New("temperature cannot be less than 273.15")
	}

	return nil
}
