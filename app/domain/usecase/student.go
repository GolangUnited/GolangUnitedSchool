package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewStudent(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *StudentUseCase {
	return &StudentUseCase{lg: lg, repo: repo}
}
