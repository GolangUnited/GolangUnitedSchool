package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentGroupUseCase struct {
	lg   *zap.Logger
	repo repository.RepositoryInterface
}

func NewStudentGroup(
	lg *zap.Logger,
	repo repository.RepositoryInterface,
) *StudentGroupUseCase {
	return &StudentGroupUseCase{lg: lg, repo: repo}
}
