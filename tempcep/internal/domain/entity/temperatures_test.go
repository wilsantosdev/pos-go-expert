package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperatures(t *testing.T) {
	t.Parallel()
	t.Run("Temperatures", func(t *testing.T) {
		t.Parallel()
		temperatures := NewTemperatures(1.0, 2.0, 3.0)
		assert.Equal(t, 1.0, temperatures.Celsius())
		assert.Equal(t, 2.0, temperatures.Farenheit())
		assert.Equal(t, 3.0, temperatures.Kelvin())
	})
}
