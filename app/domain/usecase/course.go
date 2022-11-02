package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseUsecase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewCourse(

	lg logger.Logger,
	repo repository.RepositoryInterface,
) *CourseUsecase {
	return &CourseUsecase{lg: lg, repo: repo}

}
