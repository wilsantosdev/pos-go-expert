package brasilapi

import (
	"io"
	"net/http"
	"strings"
)

const (
	BRASIL_API = "https://brasilapi.com.br/api/cep/v1/{cep}"
)

type brasilApi struct {
	channel chan string
}

func NewBrasilApi(channel *chan string) *brasilApi {
	return &brasilApi{
		channel: *channel,
	}
}

func (a brasilApi) makeUrl(cep string) string {
	return strings.ReplaceAll(BRASIL_API, "{cep}", cep)
}

func (a brasilApi) GetCEP(cep string) {
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
