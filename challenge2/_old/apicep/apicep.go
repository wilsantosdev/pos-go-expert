package apicep

import (
	"io"
	"net/http"
	"strings"
)

const (
	API_CEP = "https://cdn.apicep.com/file/apicep/{cep}.json"
)

type apiCep struct {
	channel chan string
}

func NewApiCep(channel chan string) *apiCep {
	return &apiCep{
		channel: channel,
	}
}

func (a apiCep) GetCEP(cep string) (string, error) {
	url := a.makeUrl(cep)
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

func (a apiCep) makeUrl(cep string) string {
	return strings.ReplaceAll(API_CEP, "{cep}", cep)
}
