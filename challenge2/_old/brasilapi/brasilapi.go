package brasilapi

import (
	"io"
	"net/http"
	"strings"
)

const (
	BRASIL_API = "https://brasilapi.com.br/api/cep/v1/{cep}"
)

type brasilAPI struct {
	channel *chan string
}

func NewBrasilAPI(channel *chan string) *brasilAPI {
	return &brasilAPI{
		channel: channel,
	}
}

func (b brasilAPI) GetCEP(cep string) (string, error) {
	url := b.makeUrl(cep)
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

func (b brasilAPI) makeUrl(cep string) string {
	return strings.ReplaceAll(BRASIL_API, "{cep}", cep)
}
