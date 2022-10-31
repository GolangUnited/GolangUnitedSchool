package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type Usecases struct {
	Course CourseUsecaseInterface
}

func InitUsecases(lg logger.Logger, repo repository.RepositoryInterface) *Usecases {
	return &Usecases{
		Course: NewCourse(lg, repo),
	}
}

type CourseUsecaseInterface interface {
	// AddCourse()
	// EditCourse()
	// EdotCourseByID(id int64) error
	// DeleteCourse() error
	// DeleteCoursebyID(id int64) error
}

type PersonUsecaseInterface interface {
}

type HomeworkUsecaseInterface interface {
}
