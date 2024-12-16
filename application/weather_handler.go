package application

import "context"

type WeatherOutput struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_k"`
}

type IWeatherHandler interface {
	GetTemperature(ctx context.Context, cep string) (*WeatherOutput, error)
}

type weatherHandler struct {
	coordinateService ICoordinateService
	weatherService    IWeatherService
}

func NewWeatherHandler(
	coordinateService ICoordinateService,
	weatherService IWeatherService,
) IWeatherHandler {
	return &weatherHandler{
		coordinateService: coordinateService,
		weatherService:    weatherService,
	}
}

func (w weatherHandler) GetTemperature(ctx context.Context, cep string) (*WeatherOutput, error) {
	// Tenta obter a latitude e logitude
	c, e := w.coordinateService.GetByCep(ctx, cep)

	if e != nil {
		return nil, e
	}

	// Tenta obter a temperatura
	t, e := w.weatherService.GetTemperature(ctx, c)

	if e != nil {
		return nil, e
	}

	return &WeatherOutput{
		Celsius:    t.Celsius(),
		Kelvin:     t.Kelvin(),
		Fahrenheit: t.Fahrenheit(),
	}, nil
}
