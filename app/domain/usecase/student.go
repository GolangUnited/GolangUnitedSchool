package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentUseCase struct {
	lg   *zap.Logger
	repo repository.StudentRepositoryInterface
}

func NewStudent(
	lg *zap.Logger,
	repo repository.StudentRepositoryInterface,
) *StudentUseCase {
	return &StudentUseCase{lg: lg, repo: repo}
}
