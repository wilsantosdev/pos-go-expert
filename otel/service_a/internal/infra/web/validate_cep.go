package web

import (
	"encoding/json"
	"net/http"
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

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)

}
