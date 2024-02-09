package services

import "testing"

func TestRequester(t *testing.T) {
	requester := NewRequester()

	if requester == nil {
		t.Error("Requester is nil")
	}

	requestResponse, err := requester.Request("http://www.google.com")

	if err != nil {
		t.Error("Error should be nil")
	}

	if requestResponse.GetUrl() != "http://www.google.com" {
		t.Error("Url should be http://www.google.com")
	}

	if requestResponse.GetStatusCode() != 200 {
		t.Error("StatusCode should be 200")
	}

	if requestResponse.GetError() != nil {
		t.Error("Error should be nil")
	}

	if requestResponse.GetRequestTime() == 0 {
		t.Error("RequestTime should not be 0")
	}

	if requestResponse.IsError() {
		t.Error("IsError should be false")
	}

	if !requestResponse.IsSuccess() {
		t.Error("IsSuccess should be true")
	}

	requestResponse, err = requester.Request("http://www.google.com/404")

	if err != nil {
		t.Error("Error should be nil")
	}

	if requestResponse.GetUrl() != "http://www.google.com/404" {
		t.Error("Url should be http://www.google.com/404")
	}

	if requestResponse.GetStatusCode() != 404 {
		t.Error("StatusCode should be 404")
	}

	if requestResponse.GetError() == nil {
		t.Error("Error should not be nil")
	}

	if requestResponse.GetRequestTime() == 0 {
		t.Error("RequestTime should not be 0")
	}

	if !requestResponse.IsError() {
		t.Error("IsError should be true")
	}

	if requestResponse.IsSuccess() {
		t.Error("IsSuccess should be false")
	}

}
