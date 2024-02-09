package services

import (
	"fmt"
	"stresser/internal/domain/entity"
)

type reporter struct {
	report entity.Report
}

func NewReporter(report entity.Report) *reporter {
	return &reporter{
		report: report,
	}
}

func (r *reporter) Report() {
	totalExecutionTime := r.report.TotalExecutionTime()
	totalResponses := r.report.TotalResponses()
	totalResponses200 := r.report.TotalResponses200()
	totalNon200Responses := r.report.TotalNon200Responses()

	fmt.Printf("Total execution time: %v\n", totalExecutionTime)
	fmt.Printf("Total Requests: %v\n", totalResponses)
	fmt.Printf("Total 200 Responses: %v\n", totalResponses200)

	for statusCode, total := range totalNon200Responses {
		fmt.Printf("Total %v Responses: %v\n", statusCode, total)
	}
}
