package application

import "context"

type WeatherOutput struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_k"`
}

// https://cep.awesomeapi.com.br/json/09130220
type IWeatherService interface {
	GetTemperature(ctx context.Context, cep string) (*WeatherOutput, error)
}

type weatherService struct {
	coordinateRepository ICoordinateRepository
	weatherRepository    IWeatherRepository
}

func NewWeatherService(
	coordinateRepository ICoordinateRepository,
	weatherRepository IWeatherRepository,
) IWeatherService {
	return &weatherService{
		coordinateRepository: coordinateRepository,
		weatherRepository:    weatherRepository,
	}
}

func (w weatherService) GetTemperature(ctx context.Context, cep string) (*WeatherOutput, error) {
	// Tenta obter a latitude e logitude
	c, e := w.coordinateRepository.GetCoordinate(ctx, cep)

	if e != nil {
		return nil, e
	}
	t, e := w.weatherRepository.GetTemperature(ctx, c)

	if e != nil {
		return nil, e
	}

	return &WeatherOutput{
		Celsius:    t.Celsius(),
		Kelvin:     t.Kelvin(),
		Fahrenheit: t.Fahrenheit(),
	}, nil
}
