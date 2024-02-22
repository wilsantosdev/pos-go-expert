package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCep(t *testing.T) {
	t.Parallel()
	t.Run("Cep", func(t *testing.T) {
		t.Parallel()
		cep, _ := NewCEP("12345678")
		assert.Equal(t, "12345678", cep.Value())
	})

	t.Run("CepInvalidFormat", func(t *testing.T) {
		t.Parallel()
		_, err := NewCEP("1234567")
		assert.Equal(t, "CEP invalid format", err.Error())
	})
}
