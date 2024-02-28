package domain

import "context"

type WeatherService interface {
	GetTemperatureByCity(city string, ctx context.Context) (celsius, farenheit, kelvin float64, err error)
}
