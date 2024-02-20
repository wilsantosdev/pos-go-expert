package service

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	customerror "tempcep/internal/domain/custom_error"
	"tempcep/internal/domain/entity"
)

const (
	VIA_CEP_URL = "https://viacep.com.br/ws/"
)

type viaCep struct{}

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

func NewViaCep() viaCep {
	return viaCep{}
}

func (v viaCep) GetCityNameByCep(cep entity.CEP) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(VIA_CEP_URL + cep.Value() + "/json/")
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

	return data.Localidade + " - " + data.Uf, nil

}
