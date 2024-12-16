package application

import "context"

type ICoordinateService interface {
	GetByCep(ctx context.Context, cep string) (*Coordinate, error)
}

type IWeatherService interface {
	GetTemperature(ctx context.Context, coordinate *Coordinate) (*Weather, error)
}
