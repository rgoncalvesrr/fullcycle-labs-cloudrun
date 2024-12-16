package application_test

import (
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/application"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/pkg/weather"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldReturnSuccess_WhenNewWeather(t *testing.T) {
	celsius := weather.Celsius(0.0)
	expectedCelsius := float64(celsius)
	expectedFahrenheit := celsius.ToFahrenheit()
	expectedKelvin := celsius.ToKelvin()

	w, e := application.NewWeather(celsius)

	assert.Nil(t, e)
	assert.NotNil(t, w)
	assert.Equal(t, expectedCelsius, w.Celsius())
	assert.Equal(t, expectedKelvin, w.Kelvin())
	assert.Equal(t, expectedFahrenheit, w.Fahrenheit())
}

func Test_ShouldThrowError_WhenNewWeather(t *testing.T) {
	w, e := application.NewWeather(-274)
	assert.Nil(t, w)
	assert.NotNil(t, e)
	assert.Equal(t, application.ErrInvalidTemperature.Error(), e.Error())
	assert.Equal(t, application.ErrInvalidTemperature, e)
}

func Test_ShouldThrowError_WhenInvalidCEP(t *testing.T) {
	w, e := application.NewWeather(-274)
	assert.Nil(t, w)
	assert.NotNil(t, e)
	assert.Equal(t, application.ErrInvalidTemperature.Error(), e.Error())
	assert.Equal(t, application.ErrInvalidTemperature, e)
}
