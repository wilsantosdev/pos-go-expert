package services

import (
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

	resp, err := http.Get(url)

	elapsed := time.Since(start)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code is not 200, it is %d", resp.StatusCode)
	}

	return entity.NewRequestResponse(url, int32(resp.StatusCode), err, elapsed), nil

}
