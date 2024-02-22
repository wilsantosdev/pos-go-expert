package service

import (
	customerror "tempcep/internal/domain/custom_error"
	"tempcep/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViaCep(t *testing.T) {
	var tests = []struct {
		name     string
		cep      string
		expected string
		err      *customerror.CEPNotFound
	}{
		{
			name:     "Cep",
			cep:      "23012006",
			expected: "Rio de Janeiro - RJ",
			err:      nil,
		},
		{
			name:     "CepInvalidFormat",
			cep:      "33012006",
			expected: "",
			err:      &customerror.CEPNotFound{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			viaCep := NewViaCep()
			cep, _ := entity.NewCEP(tt.cep)
			city, err := viaCep.GetCityNameByCep(cep)
			assert.Equal(t, tt.expected, city)
			if tt.err == nil {
				assert.Nil(t, err)
			}

			if tt.err != nil {
				assert.Equal(t, *tt.err, err)
			}
		})
	}

}
