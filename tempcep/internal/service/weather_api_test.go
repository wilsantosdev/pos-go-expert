package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherAPI(t *testing.T) {
	var tests = []struct {
		city string
	}{
		{
			city: "Rio de Janeiro - Brazil",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.city, func(t *testing.T) {
			t.Parallel()
			weatherApi := NewWeatherApi("281d4f6003fd4208a4823348242002")
			celsius, farenheit, kelvin, err := weatherApi.GetTemperatureByCity(tt.city)
			assert.Greater(t, celsius, 0.0)
			assert.Greater(t, farenheit, 0.0)
			assert.Greater(t, kelvin, 0.0)
			assert.Nil(t, err)
		})
	}

}
