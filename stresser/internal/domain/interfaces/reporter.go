package interfaces

import "stresser/internal/domain/entity"

type Reporter interface {
	Report(report entity.Report)
}
