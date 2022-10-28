package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type ProjectUseCase struct {
	lg   *zap.Logger
	repo repository.ProjectRepositoryInterface
}

func NewProject(lg *zap.Logger, repo repository.ProjectRepositoryInterface) *ProjectUseCase {
	return &ProjectUseCase{
		lg:   lg,
		repo: repo,
	}
}
