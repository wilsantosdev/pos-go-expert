package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	customerror "tempcep/internal/domain/custom_error"
	"tempcep/internal/usecase"
	"tempcep/internal/web/response"

	"github.com/go-chi/chi"
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
	temperatures, err := c.usecase.Execute(cep)

	switch err {
	case nil:
		resp := response.NewCepTempResponse(temperatures.Celsius(), temperatures.Farenheit(), temperatures.Kelvin())

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	case customerror.CEPInvalidFormat{}:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))

	case customerror.CEPNotFound{}:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))

	default:
		fmt.Println(err)

	}

}
