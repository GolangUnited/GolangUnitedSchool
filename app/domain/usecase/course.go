package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewCourse(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseUseCase {
	return &CourseUseCase{lg: lg, repo: repo}
}
