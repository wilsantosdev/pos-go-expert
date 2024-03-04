package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	customerror "servicea/internal/domain/custom_error"
	"servicea/internal/domain/entity"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type serviceB struct {
	serviceBURL string
	tracer      trace.Tracer
}

func NewServiceB(serviceBURL string, tracer trace.Tracer) entity.ServiceB {
	return serviceB{
		serviceBURL: serviceBURL,
		tracer:      tracer,
	}
}

func (s serviceB) GetCEPTemp(cep string, ctx context.Context) (*entity.ServiceBResponse, error) {

	ctx, span := s.tracer.Start(ctx, "get-cep-temp")
	defer span.End()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequestWithContext(ctx, "GET", s.serviceBURL+cep, nil)
	if err != nil {
		return nil, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, customerror.CEPNotFound{}
	}

	var serviceBResponse entity.ServiceBResponse
	err = json.NewDecoder(resp.Body).Decode(&serviceBResponse)
	if err != nil {
		return nil, err
	}

	return &serviceBResponse, nil
}
