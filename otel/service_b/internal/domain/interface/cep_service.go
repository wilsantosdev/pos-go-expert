package domain

import (
	"context"
	"serviceb/internal/domain/entity"
)

type CepService interface {
	GetCityNameByCep(cep entity.CEP, ctx context.Context) (string, error)
}
