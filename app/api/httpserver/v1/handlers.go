package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type Handlers struct {
	Course *CourseHandlers
	Person *PersonHandlers
}

func InitHandlers(lg logger.Logger, u usecase.Usecases) *Handlers {
	return &Handlers{
		Course: NewCourseHandler(lg, u.Course),
		Person: NewPersonHandler(lg),
	}
}
