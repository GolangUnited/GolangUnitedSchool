package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type CourseUsecase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewCourse(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseUsecase {
	return &CourseUsecase{lg: lg, repo: repo}
}
