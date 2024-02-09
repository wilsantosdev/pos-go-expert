package interfaces

import "stresser/internal/domain/entity"

type Looper interface {
	Loop(concurrent int, requests int, url string) (entity.Report, error)
}
