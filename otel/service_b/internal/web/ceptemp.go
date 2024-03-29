package web

import (
	"encoding/json"
	"net/http"
	customerror "serviceb/internal/domain/custom_error"
	"serviceb/internal/usecase"
	"serviceb/internal/web/response"

	"github.com/go-chi/chi"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type cepTempController struct {
	usecase usecase.CepTemp
}

func NewCepTempController(usecase usecase.CepTemp) *cepTempController {
	return &cepTempController{
		usecase: usecase,
	}
}

func (c *cepTempController) Handler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	carrier := propagation.HeaderCarrier(r.Header)
	c.usecase.SetCarrier(carrier)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	c.usecase.SetContext(ctx)

	temperatures, err := c.usecase.Execute(cep)

	if _, ok := err.(customerror.CEPNotFound); ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	resp := response.NewCepTempResponse(temperatures.City(), temperatures.Celsius(), temperatures.Farenheit(), temperatures.Kelvin())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
