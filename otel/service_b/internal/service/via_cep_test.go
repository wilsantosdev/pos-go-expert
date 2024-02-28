package service

import (
	"context"
	customerror "serviceb/internal/domain/custom_error"
	"serviceb/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
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
			expected: "Rio de Janeiro",
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
		tracer := otel.Tracer("service-tracer")
		viaCep := NewViaCep(tracer)
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cep, _ := entity.NewCEP(tt.cep)
			city, err := viaCep.GetCityNameByCep(cep, context.Background())
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
