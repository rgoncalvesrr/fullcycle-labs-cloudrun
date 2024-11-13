package application_test

import (
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWeather(t *testing.T) {
	w, e := application.NewWeather(0)
	assert.Nil(t, e)
	assert.NotNil(t, w)
	assert.Equal(t, 0.0, w.Celsius())
	assert.Equal(t, 273.15, w.Kelvin())
	assert.Equal(t, 32.0, w.Fahrenheit())
}

func TestNewWeatherShouldThrowError(t *testing.T) {
	w, e := application.NewWeather(-274)
	assert.Nil(t, w)
	assert.NotNil(t, e)
	assert.Equal(t, "temperature cannot be less than 273.15", e.Error())
}
