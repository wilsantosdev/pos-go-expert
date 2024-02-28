package entity

import "context"

type ServiceBResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TmepK float64 `json:"temp_K"`
}

type ServiceB interface {
	GetCEPTemp(cep string, ctx context.Context) (*ServiceBResponse, error)
}
