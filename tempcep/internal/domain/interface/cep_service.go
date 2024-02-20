package domain

import "tempcep/internal/domain/entity"

type CepService interface {
	GetCityNameByCep(cep entity.CEP) (string, error)
}
