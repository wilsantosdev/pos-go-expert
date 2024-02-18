package services

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"stresser/internal/domain/entity"
	"time"
)

type requester struct{}

func NewRequester() *requester {
	return &requester{}
}

func (r *requester) Request(url string) (*entity.RequestResponse, error) {

	start := time.Now()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	elapsed := time.Since(start)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code is not 200, it is %d", resp.StatusCode)
	}

	return entity.NewRequestResponse(url, int32(resp.StatusCode), err, elapsed), nil

}
