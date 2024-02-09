package entity

import "time"

type Report struct {
	responses []RequestResponse
}

func NewReport() *Report {
	return &Report{}
}

func (r *Report) AddResponse(response RequestResponse) {
	r.responses = append(r.responses, response)
}

func (r *Report) TotalExecutionTime() time.Duration {
	var totalExecutionTime int64
	for _, response := range r.responses {
		totalExecutionTime += int64(response.GetRequestTime())
	}
	return time.Duration(totalExecutionTime)
}

func (r *Report) TotalResponses() int32 {
	return int32(len(r.responses))
}

func (r *Report) TotalResponses200() int32 {
	var totalResponses200 int32
	for _, response := range r.responses {
		if response.GetStatusCode() == 200 {
			totalResponses200++
		}
	}
	return totalResponses200
}

func (r *Report) TotalNon200Responses() map[int32]int32 {
	totalNon200Responses := make(map[int32]int32)
	for _, response := range r.responses {
		if response.GetStatusCode() != 200 {
			totalNon200Responses[response.GetStatusCode()]++
		}
	}
	return totalNon200Responses
}
