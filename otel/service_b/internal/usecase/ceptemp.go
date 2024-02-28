package usecase

import (
	"context"
	"serviceb/internal/domain/entity"
	domain "serviceb/internal/domain/interface"

	"go.opentelemetry.io/otel/propagation"
)

type CepTemp struct {
	cepService     domain.CepService
	weatherService domain.WeatherService
	carrier        propagation.HeaderCarrier
	ctx            context.Context
}

func (u *CepTemp) SetCarrier(carrier propagation.HeaderCarrier) {
	u.carrier = carrier
}

func (u *CepTemp) SetContext(ctx context.Context) {
	u.ctx = ctx
}

func (u CepTemp) Execute(requestCEP string) (*entity.Temperatures, error) {
	cep, err := entity.NewCEP(requestCEP)
	if err != nil {
		return nil, err
	}

	// u.cepService.SetCarrier(u.carrier)
	city, err := u.cepService.GetCityNameByCep(cep, u.ctx)
	if err != nil {
		return nil, err
	}

	celsiusTemp, farenheitTemp, kelvinTemp, err := u.weatherService.GetTemperatureByCity(city, u.ctx)

	if err != nil {
		return nil, err
	}

	var temperatures = entity.NewTemperatures(city, celsiusTemp, farenheitTemp, kelvinTemp)

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
