package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseStatusUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewCourseStatus(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseStatusUseCase {
	return &CourseStatusUseCase{lg: lg, repo: repo}
}
