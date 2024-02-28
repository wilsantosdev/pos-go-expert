package entity

import (
	"regexp"
	customerror "serviceb/internal/domain/custom_error"
)

type CEP struct {
	value string
}

func NewCEP(value string) (CEP, error) {

	if !regexp.MustCompile(`^[0-9]{8}$`).MatchString(value) {
		return CEP{}, customerror.CEPInvalidFormat{}
	}

	return CEP{value: value}, nil
}

func (c CEP) Value() string {
	return c.value
}
