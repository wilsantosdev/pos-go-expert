package entity

import (
	"net/http"
	"testing"
	"time"
)

func TestReport(t *testing.T) {

	report := NewReport()

	if report == nil {
		t.Error("Report is nil")
	}

	if report.TotalResponses() != 0 {
		t.Error("TotalResponses should be 0")
	}

	if report.TotalResponses200() != 0 {
		t.Error("TotalResponses200 should be 0")
	}

	if len(report.TotalNon200Responses()) != 0 {
		t.Error("TotalNon200Responses should be empty")
	}

	requestResponse := NewRequestResponse("http://localhost", http.StatusOK, nil, time.Duration(100*time.Millisecond))

	report.AddResponse(*requestResponse)

	if report.TotalResponses() != 1 {
		t.Error("TotalResponses should be 1")
	}

	if report.TotalResponses200() != 1 {
		t.Error("TotalResponses200 should be 1")
	}

	if len(report.TotalNon200Responses()) != 0 {
		t.Error("TotalNon200Responses should be empty")
	}

	requestResponse = NewRequestResponse("http://localhost", http.StatusNotFound, nil, time.Duration(100*time.Millisecond))

	report.AddResponse(*requestResponse)

	if report.TotalResponses() != 2 {
		t.Error("TotalResponses should be 2")
	}

	if report.TotalResponses200() != 1 {
		t.Error("TotalResponses200 should be 1")
	}

	if len(report.TotalNon200Responses()) != 1 {
		t.Error("TotalNon200Responses should have 1 element")
	}

	if report.TotalExecutionTime() != time.Duration(200*time.Millisecond) {
		t.Error("TotalExecutionTime should be 200ms but got ", report.TotalExecutionTime())
	}

}
