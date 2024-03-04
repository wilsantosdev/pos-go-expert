package web

import (
	"encoding/json"
	"net/http"
	customerror "servicea/internal/domain/custom_error"
	"servicea/internal/usecase"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type validateCep struct {
	usecase usecase.ValidateCep
}

func NewValidateCep(usecase usecase.ValidateCep) validateCep {
	return validateCep{
		usecase: usecase,
	}
}

func (vc validateCep) Handler(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	var input usecase.ValidateCepInput
	err := json.NewDecoder(r.Body).Decode(&input)

	var out *usecase.ValidateCepOutput

	if err == nil {
		out, err = vc.usecase.Handler(input, ctx)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, ok := err.(customerror.CEPNotFound); ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := err.(customerror.CEPInvalidFormat); ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)

}
