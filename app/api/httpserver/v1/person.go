package v1

import "go.uber.org/zap"

type PersonHandlers struct {
	log *zap.SugaredLogger
}

func NewPersonHandlers(log *zap.SugaredLogger) *PersonHandlers {
	return &PersonHandlers{
		log: log,
	}
}
