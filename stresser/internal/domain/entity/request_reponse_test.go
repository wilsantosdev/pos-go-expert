package entity

import (
	"errors"
	"testing"
)

func TestRequestResponse(t *testing.T) {
	RequestResponse := NewRequestResponse("http://localhost", 200, nil, 100)

	if RequestResponse.GetUrl() != "http://localhost" {
		t.Error("Url should be http://localhost")
	}

	if RequestResponse.GetStatusCode() != 200 {
		t.Error("StatusCode should be 200")
	}

	if RequestResponse.GetError() != nil {
		t.Error("Error should be nil")
	}

	if RequestResponse.GetRequestTime() != 100 {
		t.Error("RequestTime should be 100")
	}

	if RequestResponse.IsError() {
		t.Error("IsError should be false")
	}

	if !RequestResponse.IsSuccess() {
		t.Error("IsSuccess should be true")
	}

	RequestResponse = NewRequestResponse("http://localhost2", 404, errors.New("Not Found"), 200)

	if RequestResponse.GetUrl() != "http://localhost2" {
		t.Error("Url should be http://localhost2")
	}

	if RequestResponse.GetStatusCode() != 404 {
		t.Error("StatusCode should be 404")
	}

	if RequestResponse.GetError() == nil {
		t.Error("Error should not be nil")
	}

	if RequestResponse.GetRequestTime() != 200 {
		t.Error("RequestTime should be 200")
	}

	if !RequestResponse.IsError() {
		t.Error("IsError should be true")
	}

	if RequestResponse.IsSuccess() {
		t.Error("IsSuccess should be false")
	}

}
