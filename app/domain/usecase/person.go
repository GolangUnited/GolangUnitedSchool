package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type PersonUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewPerson(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *CourseUseCase {
	return &CourseUseCase{lg: lg, repo: repo}
}
