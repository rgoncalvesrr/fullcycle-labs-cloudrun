package application

type IWeatherService interface {
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
