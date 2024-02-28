package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	customerror "serviceb/internal/domain/custom_error"
	"serviceb/internal/domain/entity"

	"go.opentelemetry.io/otel/trace"
)

const (
	VIA_CEP_URL = "https://viacep.com.br/ws/%v/json/"
)

type viaCep struct {
	tracer trace.Tracer
}

type viaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep(tracer trace.Tracer) *viaCep {
	return &viaCep{
		tracer: tracer,
	}
}

func (v viaCep) GetCityNameByCep(cep entity.CEP, ctx context.Context) (string, error) {

	url := fmt.Sprintf(VIA_CEP_URL, cep.Value())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	_, span := v.tracer.Start(ctx, "get-city-name")

	defer span.End()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data viaCepResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", customerror.CEPNotFound{}
	}

	return data.Localidade, nil

}
