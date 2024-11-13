package application

import "context"

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type ICoordinateRepository interface {
	GetCoordinate(ctx context.Context, cep string) (*Coordinate, error)
}

func NewCoordinate(lat, lng float64) *Coordinate {
	return &Coordinate{
		Latitude:  lat,
		Longitude: lng,
	}
}
