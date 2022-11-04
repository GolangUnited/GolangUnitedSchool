package group

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type StudentGroupUseCase struct {
	lg   logger.Logger
	repo repository.RepositoryInterface
}

func NewStudentGroup(
	lg logger.Logger,
	repo repository.RepositoryInterface,
) *StudentGroupUseCase {
	return &StudentGroupUseCase{lg: lg, repo: repo}
}
