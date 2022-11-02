package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudent(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentUseCase {
	return &StudentUseCase{lg: lg, repo: repo}
}
