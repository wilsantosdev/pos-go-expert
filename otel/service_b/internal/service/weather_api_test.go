package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
)

func TestWeatherAPI(t *testing.T) {
	var tests = []struct {
		city string
	}{
		{
			city: "Rio de Janeiro - Brazil",
		},
	}

	tracer := otel.Tracer("service-tracer")
	weatherApi := NewWeatherApi("281d4f6003fd4208a4823348242002", tracer)
	for _, tt := range tests {
		tt := tt
		t.Run(tt.city, func(t *testing.T) {
			t.Parallel()
			celsius, farenheit, kelvin, err := weatherApi.GetTemperatureByCity(tt.city, context.Background())
			assert.Greater(t, celsius, 0.0)
			assert.Greater(t, farenheit, 0.0)
			assert.Greater(t, kelvin, 0.0)
			assert.Nil(t, err)
		})
	}

}
