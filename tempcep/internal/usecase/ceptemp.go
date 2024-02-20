package usecase

import (
	"tempcep/internal/domain/entity"
	domain "tempcep/internal/domain/interface"
)

type CepTemp struct {
	cepService     domain.CepService
	weatherService domain.WeatherService
}

func (u CepTemp) Execute(requestCEP string) (*entity.Temperatures, error) {
	cep, err := entity.NewCEP(requestCEP)
	if err != nil {
		return nil, err
	}

	city, err := u.cepService.GetCityNameByCep(cep)
	if err != nil {
		return nil, err
	}

	celsiusTemp, farenheitTemp, kelvinTemp, err := u.weatherService.GetTemperatureByCity(city)

	if err != nil {
		return nil, err
	}

	var temperatures = entity.NewTemperatures(celsiusTemp, farenheitTemp, kelvinTemp)

	return &temperatures, nil

}

func NewCepTemp(
	cepService domain.CepService,
	weatherService domain.WeatherService,
) CepTemp {
	return CepTemp{
		cepService:     cepService,
		weatherService: weatherService,
	}
}
