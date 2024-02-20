package domain

type WeatherService interface {
	GetTemperatureByCity(city string) (celsius, farenheit, kelvin float64, err error)
}
