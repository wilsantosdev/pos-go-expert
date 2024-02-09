package interfaces

import "stresser/internal/domain/entity"

type Requester interface {
	Request(url string) (*entity.RequestResponse, error)
}
