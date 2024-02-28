package usecase

import (
	"context"
	"servicea/internal/domain/entity"
)

type ValidateCepInput struct {
	Cep string `json:"cep"`
}

type ValidateCepOutput struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TmepK float64 `json:"temp_K"`
}

type ValidateCep struct {
	serviceB entity.ServiceB
}

func NewValidateCep(serviceB entity.ServiceB) ValidateCep {
	return ValidateCep{
		serviceB: serviceB,
	}
}

func (vc ValidateCep) Handler(input ValidateCepInput, ctx context.Context) (*ValidateCepOutput, error) {
	cep, err := vc.validate(input)
	if err != nil {
		return nil, err
	}

	serviceBResponse, err := vc.serviceB.GetCEPTemp(cep.Value(), ctx)
	if err != nil {
		return nil, err
	}

	return &ValidateCepOutput{
		City:  serviceBResponse.City,
		TempC: serviceBResponse.TempC,
		TempF: serviceBResponse.TempF,
		TmepK: serviceBResponse.TmepK,
	}, nil
}

func (vc ValidateCep) validate(input ValidateCepInput) (cep entity.CEP, err error) {
	cep, err = entity.NewCEP(input.Cep)
	return
}
