package services

import (
	"stresser/internal/domain/entity"
	"testing"
	"time"
)

type mockRequester struct{}

func newMockRequester() *mockRequester {
	return &mockRequester{}
}

func (r *mockRequester) Request(url string) (*entity.RequestResponse, error) {
	requestResponse := entity.NewRequestResponse(url, 200, nil, time.Duration(50*time.Millisecond))
	return requestResponse, nil
}

func TestLopper(t *testing.T) {

	looper := NewLooper(newMockRequester())

	report, err := looper.Loop(10, 10, "http://www.google.com")

	if err != nil {
		t.Error("Error should be nil")
	}

	if report.TotalExecutionTime() != time.Duration(500*time.Millisecond) {
		t.Error("Total execution time should be 500ms and not ", report.TotalExecutionTime())
	}

	if report.TotalResponses() != 10 {
		t.Error("Total responses should be 10")
	}

	if report.TotalResponses200() != 10 {
		t.Error("Total 200 responses should be 10")
	}

	if len(report.TotalNon200Responses()) != 0 {
		t.Error("Total non 200 responses should be 0")
	}

}
