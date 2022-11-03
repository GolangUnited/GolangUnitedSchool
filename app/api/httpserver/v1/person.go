package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type PersonHandlers struct {
	log logger.Logger
}

func NewPersonHandler(log logger.Logger) *PersonHandlers {
	return &PersonHandlers{
		log: log,
	}
}
