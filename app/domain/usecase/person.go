package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type PersonUseCase struct {
	lg   *zap.Logger
	repo repository.CourseRepositoryInterface
}

func NewPerson(
	lg *zap.Logger,
	repo repository.CourseRepositoryInterface,
) *CourseUseCase {
	return &CourseUseCase{lg: lg, repo: repo}
}
