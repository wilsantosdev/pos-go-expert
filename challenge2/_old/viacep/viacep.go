package viacep

import (
	"io"
	"net/http"
	"strings"
)

const (
	VIA_CEP = "http://viacep.com.br/ws/{cep}/json/"
)

type viaCep struct {
	channel *chan string
}

func NewViaCep(channel *chan string) *viaCep {
	return &viaCep{
		channel: channel,
	}
}

func (v viaCep) GetCEP(cep string) (string, error) {
	url := v.makeUrl(cep)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(response), nil

}

func (v viaCep) makeUrl(cep string) string {
	return strings.ReplaceAll(VIA_CEP, "{cep}", cep)
}
