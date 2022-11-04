package course

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseStatusUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewCourseStatus(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *CourseStatusUseCase {
	return &CourseStatusUseCase{lg: lg, repo: repo}
}
