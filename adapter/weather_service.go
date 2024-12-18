package adapter

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/application"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/configs"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/pkg/weather"
)

type WeatherApiOutput struct {
	Current struct {
		Celsius float64 `json:"temp_c"`
	}
}

type weatherServiceAdapter struct {
	cfg *configs.Config
}

func NewWeatherServiceAdapter(cfg *configs.Config) application.IWeatherService {
	return &weatherServiceAdapter{cfg: cfg}
}

func (w *weatherServiceAdapter) GetTemperature(ctx context.Context, coordinate *application.Coordinate) (*application.Weather, error) {
	client := resty.New()
	r, e := client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"key": w.cfg.WeatherApiKey,
			"q":   fmt.Sprintf("%s,%s", coordinate.Latitude, coordinate.Longitude),
		}).
		SetResult(&WeatherApiOutput{}).
		Get(w.cfg.WeatherApiUrl)
	if e != nil {
		return nil, e
	}

	result, e := application.NewWeather(weather.Celsius(r.Result().(*WeatherApiOutput).Current.Celsius))

	if e != nil {
		return nil, e
	}

	return result, nil
}
