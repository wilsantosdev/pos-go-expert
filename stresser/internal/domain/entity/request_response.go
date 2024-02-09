package entity

import "time"

type RequestResponse struct {
	url         string
	statusCode  int32
	error       error
	requestTime time.Duration
}

func NewRequestResponse(url string, statusCode int32, err error, requestTime time.Duration) *RequestResponse {
	return &RequestResponse{
		url:         url,
		statusCode:  statusCode,
		error:       err,
		requestTime: requestTime,
	}
}

func (r *RequestResponse) GetUrl() string {
	return r.url
}

func (r *RequestResponse) GetStatusCode() int32 {
	return r.statusCode
}

func (r *RequestResponse) GetError() error {
	return r.error
}

func (r *RequestResponse) GetRequestTime() time.Duration {
	return r.requestTime
}

func (r *RequestResponse) IsError() bool {
	return r.error != nil
}

func (r *RequestResponse) IsSuccess() bool {
	return r.statusCode == 200 && r.error == nil
}
