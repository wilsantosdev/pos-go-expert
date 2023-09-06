package viacep

import (
	"io"
	"net/http"
	"strings"
)

const (
	VIA_CEP = "https://viacep.com.br/ws/{cep}/json/"
)

type viaCep struct {
	channel chan string
}

func NewViaCep(channel *chan string) *viaCep {
	return &viaCep{
		channel: *channel,
	}
}

func (a viaCep) makeUrl(cep string) string {
	return strings.ReplaceAll(VIA_CEP, "{cep}", cep)
}

func (a viaCep) GetCEP(cep string) {
	url := a.makeUrl(cep)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	a.channel <- string(response)

}
